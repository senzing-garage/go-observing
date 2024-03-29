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

// GrpcServerImpl is the default implementation of the GrpcServer interface.
type GrpcServerImpl struct {
	observerpb.UnimplementedObserverServer
	Port          int
	ServerOptions []grpc.ServerOption
	Subject       subject.Subject
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The Serve method starts the gRPC server.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *GrpcServerImpl) Serve(ctx context.Context) error {

	// Set up socket listener.

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", subject.Port))
	if err != nil {
		log.Printf("Port: %d; Error: %v\n", subject.Port, err)
		return err
	}
	log.Printf("Observer gRPC service running on port: %d\n", subject.Port)

	// Create server.

	aGrpcServer := grpc.NewServer(subject.ServerOptions...)
	observerpb.RegisterObserverServer(aGrpcServer, subject)

	// Enable reflection.

	reflection.Register(aGrpcServer)

	// Run server.

	err = aGrpcServer.Serve(listener)
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
func (subject *GrpcServerImpl) UpdateObserver(ctx context.Context, request *observerpb.UpdateObserverRequest) (*observerpb.UpdateObserverResponse, error) {
	var err error = nil
	if subject.Subject != nil {
		err = subject.Subject.NotifyObservers(ctx, request.GetMessage())
	}
	response := observerpb.UpdateObserverResponse{}
	return &response, err
}
