package observer

import (
	"context"
	"log"

	"github.com/senzing-garage/go-observing/observerpb"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// GrpcObserver sends the observed message to a Grpc server.
type GrpcObserver struct {
	GrpcClient observerpb.ObserverClient
	ID         string
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetObserverID method returns the unique identifier of the observer.
Use by the subject to manage the list of Observers.

Input
  - ctx: A context to control lifecycle.
*/
func (observer *GrpcObserver) GetObserverID(ctx context.Context) string {
	_ = ctx

	return observer.ID
}

/*
The UpdateObserver method processes the message sent by the Subject.
The subject invokes UpdateObserver as a goroutine.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (observer *GrpcObserver) UpdateObserver(ctx context.Context, message string) {
	if observer.GrpcClient != nil {
		request := observerpb.UpdateObserverRequest{
			Message: message,
		}

		_, err := observer.GrpcClient.UpdateObserver(ctx, &request)
		if err != nil {
			log.Printf("Observer: %s;  Message: %s; Error: %v\n", observer.ID, message, err)
		}
	}
}
