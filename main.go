/*
 */
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/senzing-garage/go-observing/grpcserver"
	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
)

const (
	port      = 8260
	sleepTime = 5
)

func main() {
	ctx := context.TODO()

	// Create a Subject.

	aSubject := subject.NewSimpleSubject()
	if aSubject.HasObservers(ctx) {
		output("Error: there shouldn't be any observers at this point.")
	}

	notifyObservers(ctx, aSubject, "Error: No observers registered, yet.")

	// Register an observer.

	anObserver1 := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := aSubject.RegisterObserver(ctx, anObserver1)
	outputError(err)
	notifyObservers(ctx, aSubject, "Message 1")

	// Register another observer.

	anObserver2 := &observer.NullObserver{
		ID: "Observer 2",
	}
	err = aSubject.RegisterObserver(ctx, anObserver2)
	outputError(err)
	notifyObservers(ctx, aSubject, "Message 2")

	// Remove observer.

	err = aSubject.UnregisterObserver(ctx, anObserver2)
	outputError(err)
	notifyObservers(ctx, aSubject, "Message 3")

	// Remove first observer.

	err = aSubject.UnregisterObserver(ctx, anObserver1)
	outputError(err)
	notifyObservers(ctx, aSubject, "Error: No observers registered, yet.")

	if aSubject.HasObservers(ctx) {
		output("Error: All observers have been removed.")
	}

	// Run an Observer gRPC service.

	err = aSubject.RegisterObserver(ctx, anObserver1)
	outputError(err)

	runGrpcServer(ctx, aSubject)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

func output(message ...any) {
	fmt.Print(message...) //nolint
}

func outputError(err error) {
	if err != nil {
		output(err)
	}
}

func notifyObservers(ctx context.Context, aSubject subject.Subject, message string) {
	if aSubject.HasObservers(ctx) {
		err := aSubject.NotifyObservers(ctx, message)
		outputError(err)
	}
}

func runGrpcServer(ctx context.Context, aSubject subject.Subject) {
	aGrpcServer := &grpcserver.SimpleGrpcServer{
		Port:    port,
		Subject: aSubject,
	}

	go func() {
		err := aGrpcServer.Serve(ctx)
		outputError(err)
	}()

	time.Sleep(sleepTime * time.Second)

	err := aGrpcServer.GracefulStop(ctx)
	outputError(err)
}
