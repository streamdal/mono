package grpcapi

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/streamdal/protos/build/go/protos"
	"github.com/streamdal/protos/build/go/protos/shared"

	"github.com/streamdal/server/util"
	"github.com/streamdal/server/validate"
)

const (
	MaxKVCommandSizeBytes = 64 * 1024 // 64KB
)

// InternalServer implements the internal GRPC API interface
type InternalServer struct {
	GRPCAPI
	// Must be implemented in order to satisfy the protos InternalServer interface
	protos.UnimplementedInternalServer
}

func (g *GRPCAPI) newInternalServer() *InternalServer {
	return &InternalServer{
		GRPCAPI: *g,
	}
}

func (s *InternalServer) sendInferSchemaPipelines(ctx context.Context, cmdCh chan *protos.Command, sessionID string) {
	// Get all audiences for this session
	audiences, err := s.Options.StoreService.GetAudiencesBySessionID(ctx, sessionID)
	if err != nil {
		s.log.Errorf("unable to get audiences by session id '%s': %v", sessionID, err)
		return
	}

	for _, aud := range audiences {
		// Create a new pipeline whose only step is an inferschema step
		attachCmd := util.GenInferSchemaPipeline(aud)

		// Inject WASM data
		if err := util.PopulateWASMFields(attachCmd.GetAttachPipeline().Pipeline, s.Options.Config.WASMDir); err != nil {
			s.log.Errorf("unable to populate WASM fields for inferschema: %v", err)
			return
		}

		cmdCh <- attachCmd
	}
}

func (s *InternalServer) Register(request *protos.RegisterRequest, server protos.Internal_RegisterServer) error {
	// validate request
	if err := validate.RegisterRequest(request); err != nil {
		return errors.Wrap(err, "invalid register request")
	}

	llog := s.log.WithFields(logrus.Fields{
		"service_name": request.ServiceName,
		"session_id":   request.SessionId,
	})

	// Store registration
	if err := s.Options.StoreService.AddRegistration(server.Context(), request); err != nil {
		return errors.Wrap(err, "unable to save registration")
	}

	// Create a new command channel
	ch, newCh := s.Options.CmdService.AddChannel(request.SessionId)

	if newCh {
		llog.Debugf("new channel created for session id '%s'", request.SessionId)
	} else {
		llog.Debugf("channel already exists for session id '%s'", request.SessionId)
	}

	var (
		shutdown bool
	)

	// Send a keepalive every tick
	ticker := time.NewTicker(1 * time.Second)

	// Broadcast registration to all nodes which will trigger handlers to push
	// an update to GetAllStream() chan (so UI knows that a change has occurred)
	if err := s.Options.BusService.BroadcastRegister(server.Context(), request); err != nil {
		return errors.Wrap(err, "unable to broadcast register")
	}

	// Send all KVs to client
	go func() {
		llog.Debug("starting initial KV sync")

		kvCommands, err := s.generateInitialKVCommands(server.Context())
		if err != nil {
			llog.Errorf("unable to generate initial kv commands: %v", err)
			return
		}

		llog.Debugf("generated '%d' kv commands", len(kvCommands))

		for _, cmd := range kvCommands {
			llog.Debugf("sending '%d' KV instructions", len(cmd.Instructions))

			ch <- &protos.Command{
				Command: &protos.Command_Kv{
					Kv: cmd,
				},
			}
		}

		llog.Debug("finished initial KV sync")
	}()

	// Send ephemeral schema inference pipeline for each announced audience
	go s.sendInferSchemaPipelines(server.Context(), ch, request.SessionId)

	// Listen for cmds from external API; forward them to connected clients
MAIN:
	for {
		select {
		case <-server.Context().Done():
			llog.Debug("register handler detected client disconnect")
			break MAIN
		case <-s.Options.ShutdownContext.Done():
			llog.Debug("register handler detected shutdown context cancellation")
			shutdown = true
			break MAIN
		case <-ticker.C:
			if err := server.Send(&protos.Command{
				Command: &protos.Command_KeepAlive{
					KeepAlive: &protos.KeepAliveCommand{},
				},
			}); err != nil {
				// TODO: If unable to send heartbeat to client X times, stop request/exit loop
				llog.WithError(err).Errorf("unable to send heartbeat for session id '%s'", request.SessionId)
				continue
			}
		case cmd := <-ch:
			if cmd == nil {
				llog.Warning("received nil cmd on cmd channel; ignoring")
				continue
			}

			llog.Debugf("received cmd on cmd channel; sending cmd to client session id '%s'", request.SessionId)

			// Send cmd to connected client
			if err := server.Send(cmd); err != nil {
				s.log.WithError(err).Error("unable to send cmd to client")

				// TODO: Retry? Ignore?
				return errors.Wrap(err, "unable to send cmd to client")
			}

			llog.Debug("sent cmd to client")
		}
	}

	if shutdown {
		llog.Debugf("register handler shutting down for req id '%s'", server.Context().Value("id"))

		// Notify client that they need to re-register because of shutdown
		return GRPCServerShutdownError
	}

	llog.Debug("client has disconnected; de-registering")

	// Remove command channel
	if ok := s.Options.CmdService.RemoveChannel(request.SessionId); ok {
		llog.Debug("removed cmd channel")
	} else {
		llog.Debug("no cmd channel found")
	}

	deregisterRequest := &protos.DeregisterRequest{
		ServiceName: request.ServiceName,
		SessionId:   request.SessionId,
	}

	llog.Debug("deleting registration from store")

	// By this point, the server context may be cancelled, so we must not rely on it
	ctx := context.Background()

	if err := s.Options.StoreService.DeleteRegistration(ctx, deregisterRequest); err != nil {
		err = errors.Wrap(err, "unable to delete registration")
		llog.Error(err)
		return err
	}

	llog.Debug("announcing de-registration")

	// Announce de-registration - the UI will still display the audience(s) but
	// they no longer will be live (ie. attached clients will decrease)
	if err := s.Options.BusService.BroadcastDeregister(ctx, deregisterRequest); err != nil {
		llog.Errorf("unable to broadcast deregister event: %s", err)
	}

	return nil
}

func (s *InternalServer) Heartbeat(ctx context.Context, req *protos.HeartbeatRequest) (*protos.StandardResponse, error) {
	if err := validate.HeartbeatRequest(req); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST,
			Message: fmt.Sprintf("invalid heartbeat req: %s", err.Error()),
		}, nil
	}

	// Refresh register key
	// This method also refreshes audience live keys
	if err := s.Options.StoreService.AddRegistration(ctx, &protos.RegisterRequest{
		ServiceName: req.ServiceName,
		SessionId:   req.SessionId,
		ClientInfo:  req.ClientInfo,
		Audiences:   req.Audiences,
	}); err != nil {
		err = errors.Wrap(err, "unable to save heartbeat")
		s.log.Error(err)

		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR,
			Message: err.Error(),
		}, nil
	}

	s.log.Debug("Saved heartbeat")

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Heartbeat received",
	}, nil
}

func (s *InternalServer) Notify(ctx context.Context, request *protos.NotifyRequest) (*protos.StandardResponse, error) {
	if err := s.Options.NotifyService.Queue(ctx, request); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR,
			Message: fmt.Sprintf("unable to queue notification: %s", err.Error()),
		}, nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Notification queued",
	}, nil
}

func (s *InternalServer) Metrics(ctx context.Context, req *protos.MetricsRequest) (*protos.StandardResponse, error) {
	if err := validate.MetricsRequest(req); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST,
			Message: fmt.Sprintf("invalid metrics req: %s", err.Error()),
		}, nil
	}

	if err := s.Options.BusService.BroadcastMetrics(ctx, req); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR,
			Message: fmt.Sprintf("unable to handle metrics request: %s", err.Error()),
		}, nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Metrics handled",
	}, nil
}

func (s *InternalServer) NewAudience(ctx context.Context, req *protos.NewAudienceRequest) (*protos.StandardResponse, error) {
	s.log.Debugf("received new audience request for session id '%s'", req.SessionId)

	if err := validate.NewAudienceRequest(req); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST,
			Message: fmt.Sprintf("invalid new audience req: %s", err.Error()),
		}, nil
	}

	if err := s.Options.StoreService.AddAudience(ctx, req); err != nil {
		s.log.Errorf("unable to save audience: %s", err.Error())

		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR,
			Message: fmt.Sprintf("unable to save audience: %s", err.Error()),
		}, nil
	}

	// Send AttachCommand to client with ephemeral inferschema pipeline
	cmdCh, isNewCh := s.Options.CmdService.AddChannel(req.SessionId)
	if isNewCh {
		s.log.Debugf("new channel created for session id '%s'", req.SessionId)
	} else {
		s.log.Debugf("channel already exists for session id '%s'", req.SessionId)
	}

	// This is context.Background() because it's ran as a gouroutine and the request
	// context may be finished by the time it eventually runs
	go s.sendInferSchemaPipelines(context.Background(), cmdCh, req.SessionId)

	// Broadcast audience creation so that we can notify UI GetAllStream clients
	if err := s.Options.BusService.BroadcastNewAudience(ctx, req); err != nil {
		s.log.Errorf("unable to broadcast new audience: %s", err.Error())
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Audience created",
	}, nil
}

// GetActiveCommands is used by SDKs to get all active commands at SDK init -
// this enables the SDKs to "resume" whatever operations that were active when
// they were last connected to the Streamdal server.
func (s *InternalServer) GetActiveCommands(
	ctx context.Context,
	req *protos.GetActiveCommandsRequest,
) (*protos.GetActiveCommandsResponse, error) {
	if req.ServiceName == "" {
		return nil, errors.New("service name is required")
	}

	active, paused, err := s.getActiveAttachPipelineCommands(ctx, req.ServiceName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get attach commands by service")
	}

	tailCommands, err := s.Options.StoreService.GetActiveTailCommandsByService(ctx, req.ServiceName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get active tails")
	}

	return &protos.GetActiveCommandsResponse{
		Active:      append(active, tailCommands...), // Include start tail commands in active commands
		Paused:      paused,
		WasmModules: util.GenerateWasmMapping(append(active, paused...)...), // Include active and paused; tail does not have steps
	}, nil
}

func (s *InternalServer) getActiveAttachPipelineCommands(
	ctx context.Context,
	serviceName string,
) ([]*protos.Command, []*protos.Command, error) {
	if serviceName == "" {
		return nil, nil, errors.New("service name is required")
	}

	attaches, err := s.Options.StoreService.GetAttachCommandsByService(ctx, serviceName)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to get attach commands by service")
	}

	pausedMap, err := s.Options.StoreService.GetPaused(ctx)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to get paused pipelines")
	}

	active := make([]*protos.Command, 0)
	paused := make([]*protos.Command, 0)

	for _, a := range attaches {
		if err := validate.AttachPipelineCommand(a.GetAttachPipeline()); err != nil {
			s.log.Warningf("invalid attach pipeline command: %s", err.Error())
			continue
		}

		// Always wipe bytes. SDK Client will handle lookup.
		for _, step := range a.GetAttachPipeline().Pipeline.Steps {
			step.XWasmBytes = nil
		}

		pausedEntry, ok := pausedMap[a.GetAttachPipeline().Pipeline.Id]
		if !ok {
			// Pipeline ID is not present in map, it is not paused
			active = append(active, a)
			continue
		}

		// Found pipeline id in paused, check if audience matches
		if util.AudienceEquals(pausedEntry.Audience, a.Audience) {
			// This is paused
			paused = append(paused, a)
			continue
		}

		// Audience dos not match, this is active
		active = append(active, a)
	}

	return active, paused, nil
}

// GetAttachCommandsByService is used by SDKs to get active attach commands at
// SDK initialization time. It is 'ByService' because that is the only thing
// an SDK knows 100% about itself at initialization time -- audiences may not
// be announced until a later .Process() call.
//
// DEPRECATED - use GetActiveCommands() instead
func (s *InternalServer) GetAttachCommandsByService(
	ctx context.Context,
	req *protos.GetAttachCommandsByServiceRequest,
) (*protos.GetAttachCommandsByServiceResponse, error) {
	active, paused, err := s.getActiveAttachPipelineCommands(ctx, req.ServiceName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get attach commands by service")
	}

	return &protos.GetAttachCommandsByServiceResponse{
		Active:      active,
		Paused:      paused,
		WasmModules: util.GenerateWasmMapping(append(active, paused...)...),
	}, nil
}

func (s *InternalServer) SendTail(srv protos.Internal_SendTailServer) error {
	llog := s.log.WithFields(logrus.Fields{
		"method":     "SendTail",
		"request_id": util.CtxRequestId(srv.Context()),
	})

	// This isn't necessary for go, but other langauge libraries, such as python
	// require a response to eventually be sent and will throw an exception if
	// one is not received
	defer func() {
		llog.Debug("sending final tail response")

		srv.SendAndClose(&protos.StandardResponse{
			Id:   util.CtxRequestId(srv.Context()),
			Code: protos.ResponseCode_RESPONSE_CODE_OK,
		})
	}()

	for {
		select {
		case <-srv.Context().Done():
			llog.Debug("detected client disconnect")
			return nil
		case <-s.Options.ShutdownContext.Done():
			llog.Debug("server shutting down, exiting stream")
			return nil
		default:
			tailResp, err := srv.Recv()
			if err != nil {
				if strings.Contains(err.Error(), io.EOF.Error()) || strings.Contains(err.Error(), context.Canceled.Error()) {
					llog.Debug("client closed stream")
					return nil
				}

				llog.Error(errors.Wrap(err, "unable to receive tail response"))
				continue
			}

			llog.Debugf("after srv.Recv(), before validation for session id '%s', tail request id '%s'",
				tailResp.SessionId, tailResp.TailRequestId)

			if err := validate.TailResponse(tailResp); err != nil {
				llog.Error(errors.Wrapf(err, "invalid tail response for session id '%s', tail request id '%s'",
					tailResp.SessionId, tailResp.TailRequestId))
				continue
			}

			llog.Debugf("after validation, before BroadcastTailResponse() for session id '%s', tail request id '%s'",
				tailResp.SessionId, tailResp.TailRequestId)

			if err := s.Options.BusService.BroadcastTailResponse(srv.Context(), tailResp); err != nil {
				llog.Error(errors.Wrapf(err, "unable to broadcast tail response for session id '%s', tail request id '%s'",
					tailResp.SessionId, tailResp.TailRequestId))

				continue
			}

			llog.Debugf("after BroadcastTailResponse() for session id '%s', tail request id '%s'",
				tailResp.SessionId, tailResp.TailRequestId)
		}
	}
}

// Will generate a batch of KVCommands that are intended to be sent to SDK
// clients upon registration
func (s *InternalServer) generateInitialKVCommands(ctx context.Context) ([]*protos.KVCommand, error) {
	// Fetch all KVs
	kvs, err := s.Options.KVService.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get all KVs")
	}

	cmds := make([]*protos.KVCommand, 0)
	instructions := make([]*protos.KVInstruction, 0)
	size := 0

	// Inject up to 64KB of instructions per KVCommand
	for _, kv := range kvs {
		if size > MaxKVCommandSizeBytes {
			// Copy instructions
			instructionsCopy := make([]*protos.KVInstruction, len(instructions))
			copy(instructionsCopy, instructions)

			// Append a new command w/ instructions
			cmds = append(cmds, &protos.KVCommand{
				Instructions: instructionsCopy,
				Overwrite:    true,
			})

			// Reset instructions
			instructions = make([]*protos.KVInstruction, 0)
			size = 0
		}

		instructions = append(instructions, &protos.KVInstruction{
			Id:                       util.GenerateUUID(),
			Action:                   shared.KVAction_KV_ACTION_CREATE,
			Object:                   kv,
			RequestedAtUnixTsNanoUtc: time.Now().UTC().UnixNano(),
		})

		size += len(kv.Key) + len(kv.Value)
	}

	// Append remainder to cmds
	cmds = append(cmds, &protos.KVCommand{
		Instructions: instructions,
	})

	s.log.Debugf("generateInitialKVCommands has generated '%d' KV commands", len(cmds))

	return cmds, nil
}

func (s *InternalServer) SendSchema(ctx context.Context, req *protos.SendSchemaRequest) (*protos.StandardResponse, error) {
	if err := validate.SendSchemaRequest(req); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST,
			Message: fmt.Sprintf("invalid request: %s", err.Error()),
		}, nil
	}

	// Get existing schema
	schema, err := s.Options.StoreService.GetSchema(ctx, req.Audience)
	if err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_BAD_REQUEST,
			Message: errors.Wrap(err, "error getting existing schema").Error(),
		}, nil
	}

	// Bump version and timestamp
	// GetSchema returns an empty schema with version 0 if it doesn't exist yet
	schema.XVersion++
	schema.JsonSchema = req.Schema.JsonSchema
	schema.XMetadata = req.Schema.XMetadata

	if req.Schema.XMetadata == nil {
		req.Schema.XMetadata = make(map[string]string)
	}

	req.Schema.XMetadata["last_updated"] = time.Now().UTC().Format(time.RFC3339)

	if err := s.Options.StoreService.AddSchema(ctx, req); err != nil {
		return &protos.StandardResponse{
			Id:      util.CtxRequestId(ctx),
			Code:    protos.ResponseCode_RESPONSE_CODE_INTERNAL_SERVER_ERROR,
			Message: fmt.Sprintf("unable to save schema: %s", err.Error()),
		}, nil
	}

	return &protos.StandardResponse{
		Id:      util.CtxRequestId(ctx),
		Code:    protos.ResponseCode_RESPONSE_CODE_OK,
		Message: "Schema received",
	}, nil
}
