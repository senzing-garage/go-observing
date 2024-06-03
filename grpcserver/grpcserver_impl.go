package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/senzing-garage/go-observing/observerpb"
	"github.com/senzing-garage/go-observing/subject"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SimpleGrpcServer is the default implementation of the GrpcServer interface.
type SimpleGrpcServer struct {
	observerpb.UnimplementedObserverServer
	Port          int
	ServerOptions []grpc.ServerOption
	Subject       subject.Subject
	server        *grpc.Server
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Serve method starts the gRPC server.

Input
  - ctx: A context to control lifecycle.
*/
func (grpcServer *SimpleGrpcServer) GracefulStop(ctx context.Context) error {
	_ = ctx
	grpcServer.server.GracefulStop()
	return nil
}

/*
The Serve method starts the gRPC server.

Input
  - ctx: A context to control lifecycle.
*/
func (grpcServer *SimpleGrpcServer) Serve(ctx context.Context) error {
	_ = ctx

	// Set up socket listener.

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcServer.Port))
	if err != nil {
		log.Printf("Port: %d; Error: %v\n", grpcServer.Port, err)
		return err
	}
	log.Printf("Observer gRPC service running on port: %d\n", grpcServer.Port)

	// Create server.

	grpcServer.server = grpc.NewServer(grpcServer.ServerOptions...)
	observerpb.RegisterObserverServer(grpcServer.server, grpcServer)

	// Enable reflection.

	reflection.Register(grpcServer.server)

	// Run server.

	err = grpcServer.server.Serve(listener)
	if err != nil {
		log.Println(err)
	}

	return err
}

/*
The UpdateObserver method is a gRPC protocol method which receives a message from a remote
observer and repeats it to local observers.

Input
  - ctx: A context to control lifecycle.
  - request: Includes the message to send to Observers.

Output
  - Empty response
  - Error
*/
func (grpcServer *SimpleGrpcServer) UpdateObserver(ctx context.Context, request *observerpb.UpdateObserverRequest) (*observerpb.UpdateObserverResponse, error) {
	_ = ctx
	var err error
	if grpcServer.Subject != nil {
		err = grpcServer.Subject.NotifyObservers(ctx, request.GetMessage())
	}
	response := observerpb.UpdateObserverResponse{}
	return &response, err
}
