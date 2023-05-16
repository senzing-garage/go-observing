/*
 */
package main

import (
	"context"
	"fmt"

	"github.com/senzing/go-observing/grpcserver"
	"github.com/senzing/go-observing/observer"
	"github.com/senzing/go-observing/subject"
)

func main() {
	ctx := context.TODO()

	// Create a Subject.

	aSubject := &subject.SubjectImpl{}

	if aSubject.HasObservers(ctx) {
		fmt.Print("Error: there shouldn't be any observers at this point.")
	}

	err := aSubject.NotifyObservers(ctx, "Error: No observers registered, yet.")
	if err != nil {
		fmt.Print(err)
	}

	// Register an observer.

	anObserver1 := &observer.ObserverNull{
		Id: "Observer 1",
	}
	err = aSubject.RegisterObserver(ctx, anObserver1)
	if err != nil {
		fmt.Print(err)
	}

	// Notify.

	err = aSubject.NotifyObservers(ctx, "Message 1")
	if err != nil {
		fmt.Print(err)
	}

	// Register another observer.

	anObserver2 := &observer.ObserverNull{
		Id: "Observer 2",
	}
	err = aSubject.RegisterObserver(ctx, anObserver2)
	if err != nil {
		fmt.Print(err)
	}

	// Notify.

	if aSubject.HasObservers(ctx) {
		err = aSubject.NotifyObservers(ctx, "Message 2")
		if err != nil {
			fmt.Print(err)
		}
	}

	// Remove observer.

	err = aSubject.UnregisterObserver(ctx, anObserver2)
	if err != nil {
		fmt.Print(err)
	}

	// Notify.

	if aSubject.HasObservers(ctx) {
		err = aSubject.NotifyObservers(ctx, "Message 3")
		if err != nil {
			fmt.Print(err)
		}
	}

	// Remove observer.

	err = aSubject.UnregisterObserver(ctx, anObserver1)
	if err != nil {
		fmt.Print(err)
	}

	// Notify.

	err = aSubject.NotifyObservers(ctx, "Error: No observers registered, yet.")
	if err != nil {
		fmt.Print(err)
	}

	if aSubject.HasObservers(ctx) {
		fmt.Print("Error: All observers have been removed.")
	}

	// Run an Observer gRPC service.

	err = aSubject.RegisterObserver(ctx, anObserver1)
	if err != nil {
		fmt.Print(err)
	}

	aGrpcServer := &grpcserver.GrpcServerImpl{
		Port:    8260,
		Subject: aSubject,
	}
	err = aGrpcServer.Serve(ctx)
	if err != nil {
		fmt.Print(err)
	}
}
