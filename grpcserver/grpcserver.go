package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/senzing-garage/go-helpers/wraperror"
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
	Port          int
	ServerOptions []grpc.ServerOption
	Subject       subject.Subject
	observerpb.UnimplementedObserverServer
	server *grpc.Server
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
	const (
		keepAliveSeconds = 30
	)

	// Set up socket listener.

	listenConfig := &net.ListenConfig{ //nolint
		KeepAlive: keepAliveSeconds * time.Second,
	}

	listener, err := listenConfig.Listen(ctx, "tcp", fmt.Sprintf(":%d", grpcServer.Port))
	if err != nil {
		log.Printf("Port: %d; Error: %v\n", grpcServer.Port, err)

		return wraperror.Errorf(err, "net.Listen on port %d", grpcServer.Port)
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

	return wraperror.Errorf(err, wraperror.NoMessage)
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
func (grpcServer *SimpleGrpcServer) UpdateObserver(
	ctx context.Context,
	request *observerpb.UpdateObserverRequest,
) (*observerpb.UpdateObserverResponse, error) {
	_ = ctx

	var err error
	if grpcServer.Subject != nil {
		err = grpcServer.Subject.NotifyObservers(ctx, request.GetMessage())
	}

	response := observerpb.UpdateObserverResponse{}

	return &response, wraperror.Errorf(err, wraperror.NoMessage)
}
