package grpcsubject

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	grpcx "github.com/senzing/go-observing/grpcx"
	"github.com/senzing/go-observing/observer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SubjectImpl is the default implementation of the Subject interface.
type SubjectImpl struct {
	grpcx.UnimplementedObserverServer
	observerList []observer.Observer
	Port         int
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var lock = sync.RWMutex{}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func contains(ctx context.Context, haystack []observer.Observer, needle observer.Observer) bool {
	needleId := needle.GetObserverId(ctx)
	for _, value := range haystack {
		if value.GetObserverId(ctx) == needleId {
			return true
		}
	}
	return false
}

func removeFromSlice(ctx context.Context, observerList []observer.Observer, observerToRemove observer.Observer) []observer.Observer {
	removeId := observerToRemove.GetObserverId(ctx)
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observer.GetObserverId(ctx) == removeId {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetObservers method returns a clone of registered Observers.
The caller will not have access to the type-struct's internal slice of observers.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *SubjectImpl) GetObservers(ctx context.Context) []observer.Observer {
	result := make([]observer.Observer, len(subject.observerList))
	copy(result, subject.observerList)
	return result
}

/*
The HasObservers method is used to determine if there are any Observers.
It can be used as a "guard" to avoid creating a message only to have
nothing observe it.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *SubjectImpl) HasObservers(ctx context.Context) bool {
	return len(subject.observerList) > 0
}

/*
The NotifyObservers method notifies all listeners of the message.
This is done asynchronously using goroutines.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (subject *SubjectImpl) NotifyObservers(ctx context.Context, message string) error {
	var err error = nil
	for _, observer := range subject.observerList {
		go observer.UpdateObserver(ctx, message)
	}
	return err
}

/*
The RegisterObserver method adds an observer.Observer to an internal list of
listeners to be notified. The listener will receive event notifications.

Input
  - ctx: A context to control lifecycle.
  - observer: A component wanting to listen to events.
*/
func (subject *SubjectImpl) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error = nil
	lock.RLock()
	defer lock.RUnlock()
	if !contains(ctx, subject.observerList, observer) {
		subject.observerList = append(subject.observerList, observer)
	}
	return err
}

/*
The Serve method starts the gRPC server.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *SubjectImpl) Serve(ctx context.Context) error {

	// Set up socket listener.

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", subject.Port))
	if err != nil {
		log.Printf("Port: %d; Error: %v\n", subject.Port, err)
		return err
	}
	log.Printf("Observer gRPC service running on port: %d\n", subject.Port)

	// Create server.

	aGrpcServer := grpc.NewServer()
	grpcx.RegisterObserverServer(aGrpcServer, subject)

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
The UnregisterObserver method removes an observer.Observer from an internal list of
listeners.  The listener will no longer receive event notifications.

Input
  - ctx: A context to control lifecycle.
  - observer: A component no longer wanting to listen to events.
*/
func (subject *SubjectImpl) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error = nil
	lock.RLock()
	defer lock.RUnlock()
	subject.observerList = removeFromSlice(ctx, subject.observerList, observer)
	return err
}

/*
The UpdateObserver method is a gRPC protocol method which receives a message from a remote
observer and repeats it to local observers.

Input
  - ctx: A context to control lifecycle.
  - request: A component no longer wanting to listen to events.

Output
  - Empty response
  - Error
*/
func (subject *SubjectImpl) UpdateObserver(ctx context.Context, request *grpcx.UpdateObserverRequest) (*grpcx.UpdateObserverResponse, error) {
	err := subject.NotifyObservers(ctx, request.GetMessage())
	response := grpcx.UpdateObserverResponse{}
	return &response, err
}
