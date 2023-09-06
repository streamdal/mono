// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_internal.proto" (package "protos", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { Internal } from "./sp_internal.ts";
import type { TailResponse } from "./sp_common.ts";
import type { ClientStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { GetAttachCommandsByServiceResponse } from "./sp_internal.ts";
import type { GetAttachCommandsByServiceRequest } from "./sp_internal.ts";
import type { MetricsRequest } from "./sp_internal.ts";
import type { NotifyRequest } from "./sp_internal.ts";
import type { HeartbeatRequest } from "./sp_internal.ts";
import type { StandardResponse } from "./sp_common.ts";
import type { NewAudienceRequest } from "./sp_internal.ts";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { Command } from "./sp_command.ts";
import type { RegisterRequest } from "./sp_internal.ts";
import type { ServerStreamingCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service protos.Internal
 */
export interface IInternalClient {
    /**
     * Initial method that an SDK should call to register itself with the server.
     * The server will use this stream to send commands to the SDK via the
     * `CommandResponse` message. Clients should continuously listen for
     * CommandResponse messages and re-establish registration if the stream gets
     * disconnected.
     *
     * @generated from protobuf rpc: Register(protos.RegisterRequest) returns (stream protos.Command);
     */
    register(input: RegisterRequest, options?: RpcOptions): ServerStreamingCall<RegisterRequest, Command>;
    /**
     * Declare a new audience that the SDK is able to accept commands for.
     * An SDK would use this method when a new audience is declared by the user
     * via `.Process()`.
     *
     * @generated from protobuf rpc: NewAudience(protos.NewAudienceRequest) returns (protos.StandardResponse);
     */
    newAudience(input: NewAudienceRequest, options?: RpcOptions): UnaryCall<NewAudienceRequest, StandardResponse>;
    /**
     * SDK is responsible for sending heartbeats to the server to let the server
     * know about active consumers and producers.
     *
     * @generated from protobuf rpc: Heartbeat(protos.HeartbeatRequest) returns (protos.StandardResponse);
     */
    heartbeat(input: HeartbeatRequest, options?: RpcOptions): UnaryCall<HeartbeatRequest, StandardResponse>;
    /**
     * Use this method when Notify condition has been triggered; the server will
     * decide on what to do about the notification.
     *
     * @generated from protobuf rpc: Notify(protos.NotifyRequest) returns (protos.StandardResponse);
     */
    notify(input: NotifyRequest, options?: RpcOptions): UnaryCall<NotifyRequest, StandardResponse>;
    /**
     * Send periodic metrics to the server
     *
     * @generated from protobuf rpc: Metrics(protos.MetricsRequest) returns (protos.StandardResponse);
     */
    metrics(input: MetricsRequest, options?: RpcOptions): UnaryCall<MetricsRequest, StandardResponse>;
    /**
     * Used to pull all pipeline configs for the service name in the SDK's constructor
     * This is needed because Register() is async
     *
     * @generated from protobuf rpc: GetAttachCommandsByService(protos.GetAttachCommandsByServiceRequest) returns (protos.GetAttachCommandsByServiceResponse);
     */
    getAttachCommandsByService(input: GetAttachCommandsByServiceRequest, options?: RpcOptions): UnaryCall<GetAttachCommandsByServiceRequest, GetAttachCommandsByServiceResponse>;
    /**
     * @generated from protobuf rpc: SendTail(stream protos.TailResponse) returns (protos.StandardResponse);
     */
    sendTail(options?: RpcOptions): ClientStreamingCall<TailResponse, StandardResponse>;
}
/**
 * @generated from protobuf service protos.Internal
 */
export class InternalClient implements IInternalClient, ServiceInfo {
    typeName = Internal.typeName;
    methods = Internal.methods;
    options = Internal.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * Initial method that an SDK should call to register itself with the server.
     * The server will use this stream to send commands to the SDK via the
     * `CommandResponse` message. Clients should continuously listen for
     * CommandResponse messages and re-establish registration if the stream gets
     * disconnected.
     *
     * @generated from protobuf rpc: Register(protos.RegisterRequest) returns (stream protos.Command);
     */
    register(input: RegisterRequest, options?: RpcOptions): ServerStreamingCall<RegisterRequest, Command> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<RegisterRequest, Command>("serverStreaming", this._transport, method, opt, input);
    }
    /**
     * Declare a new audience that the SDK is able to accept commands for.
     * An SDK would use this method when a new audience is declared by the user
     * via `.Process()`.
     *
     * @generated from protobuf rpc: NewAudience(protos.NewAudienceRequest) returns (protos.StandardResponse);
     */
    newAudience(input: NewAudienceRequest, options?: RpcOptions): UnaryCall<NewAudienceRequest, StandardResponse> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<NewAudienceRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * SDK is responsible for sending heartbeats to the server to let the server
     * know about active consumers and producers.
     *
     * @generated from protobuf rpc: Heartbeat(protos.HeartbeatRequest) returns (protos.StandardResponse);
     */
    heartbeat(input: HeartbeatRequest, options?: RpcOptions): UnaryCall<HeartbeatRequest, StandardResponse> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<HeartbeatRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Use this method when Notify condition has been triggered; the server will
     * decide on what to do about the notification.
     *
     * @generated from protobuf rpc: Notify(protos.NotifyRequest) returns (protos.StandardResponse);
     */
    notify(input: NotifyRequest, options?: RpcOptions): UnaryCall<NotifyRequest, StandardResponse> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<NotifyRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Send periodic metrics to the server
     *
     * @generated from protobuf rpc: Metrics(protos.MetricsRequest) returns (protos.StandardResponse);
     */
    metrics(input: MetricsRequest, options?: RpcOptions): UnaryCall<MetricsRequest, StandardResponse> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<MetricsRequest, StandardResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * Used to pull all pipeline configs for the service name in the SDK's constructor
     * This is needed because Register() is async
     *
     * @generated from protobuf rpc: GetAttachCommandsByService(protos.GetAttachCommandsByServiceRequest) returns (protos.GetAttachCommandsByServiceResponse);
     */
    getAttachCommandsByService(input: GetAttachCommandsByServiceRequest, options?: RpcOptions): UnaryCall<GetAttachCommandsByServiceRequest, GetAttachCommandsByServiceResponse> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<GetAttachCommandsByServiceRequest, GetAttachCommandsByServiceResponse>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SendTail(stream protos.TailResponse) returns (protos.StandardResponse);
     */
    sendTail(options?: RpcOptions): ClientStreamingCall<TailResponse, StandardResponse> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<TailResponse, StandardResponse>("clientStreaming", this._transport, method, opt);
    }
}
