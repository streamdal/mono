package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/streamdal/mono/libs/protos/build/go/protos"
)

func main() {
	// dial server
	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	// create stream
	client := protos.NewInternalClient(conn)

	session_id, _ := uuid.NewV4()

	// Generate register request
	in := &protos.RegisterRequest{
		ServiceName: "TestService",
		DryRun:      false,
		SessionId:   session_id.String(),
		ClientInfo: &protos.ClientInfo{
			ClientType:     2,
			LibraryName:    "demo-client",
			LibraryVersion: "1.0",
			Language:       "go",
			Arch:           runtime.GOARCH,
			Os:             runtime.GOOS,
		},
	}

	md := metadata.New(map[string]string{"auth-token": "1234"})

	// Send register request
	stream, err := client.Register(metadata.NewOutgoingContext(context.Background(), md), in)
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}

	done := make(chan bool)

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true //means stream is finished
				return
			}
			if err != nil {
				log.Fatalf("cannot receive %v", err)
			}
			log.Printf("Resp received: %s", resp.GetCommand())
		}
	}()

	fmt.Println("Server won't EOF; press Ctrl+C to stop client")

	<-done
	log.Printf("finished")
}

//func main() {
//	// dial server
//	conn, err := grpc.Dial(":8082", grpc.WithInsecure())
//	if err != nil {
//		log.Fatalf("can not connect with server %v", err)
//	}
//
//	// create stream
//	client := protos.NewInternalClient(conn)
//
//	// Generate register request
//	in := &protos.HeartbeatRequest{}
//
//	// Send register request
//	resp, err := client.Heartbeat(context.Background(), in)
//	if err != nil {
//		log.Fatalf("open stream error %v", err)
//	}
//
//	fmt.Printf("Received response: %+v\n", resp)
//
//	log.Printf("finished")
//}
