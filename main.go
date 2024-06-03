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

func printError(err error) {
	if err != nil {
		fmt.Print(err)
	}
}

func main() {
	ctx := context.TODO()

	// Create a Subject.

	aSubject := &subject.SimpleSubject{}

	if aSubject.HasObservers(ctx) {
		fmt.Print("Error: there shouldn't be any observers at this point.")
	}

	err := aSubject.NotifyObservers(ctx, "Error: No observers registered, yet.")
	printError(err)

	// Register an observer.

	anObserver1 := &observer.NullObserver{
		ID: "Observer 1",
	}
	err = aSubject.RegisterObserver(ctx, anObserver1)
	printError(err)

	// Notify.

	err = aSubject.NotifyObservers(ctx, "Message 1")
	printError(err)

	// Register another observer.

	anObserver2 := &observer.NullObserver{
		ID: "Observer 2",
	}
	err = aSubject.RegisterObserver(ctx, anObserver2)
	printError(err)

	// Notify.

	if aSubject.HasObservers(ctx) {
		err = aSubject.NotifyObservers(ctx, "Message 2")
		printError(err)
	}

	// Remove observer.

	err = aSubject.UnregisterObserver(ctx, anObserver2)
	printError(err)

	// Notify.

	if aSubject.HasObservers(ctx) {
		err = aSubject.NotifyObservers(ctx, "Message 3")
		printError(err)
	}

	// Remove observer.

	err = aSubject.UnregisterObserver(ctx, anObserver1)
	printError(err)

	// Notify.

	err = aSubject.NotifyObservers(ctx, "Error: No observers registered, yet.")
	printError(err)

	if aSubject.HasObservers(ctx) {
		fmt.Print("Error: All observers have been removed.")
	}

	// Run an Observer gRPC service.

	err = aSubject.RegisterObserver(ctx, anObserver1)
	printError(err)

	aGrpcServer := &grpcserver.SimpleGrpcServer{
		Port:    8260,
		Subject: aSubject,
	}

	go func() {
		err = aGrpcServer.Serve(ctx)
		printError(err)
	}()

	time.Sleep(5 * time.Second)

	err = aGrpcServer.GracefulStop(ctx)
	printError(err)
}
