/*
 */
package main

import (
	"context"
	"fmt"

	"github.com/senzing/go-observing/grpcsubject"
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

	// Run a gRPC server

	aGrpcSubject := &grpcsubject.SubjectImpl{
		Port: 4100,
	}
	err = aGrpcSubject.RegisterObserver(ctx, anObserver1)
	if err != nil {
		fmt.Print(err)
	}
	aGrpcSubject.Serve(ctx)

	// Give time to allow Observers to print.

	// sleepDuration := 2 * time.Second
	// fmt.Printf("Sleeping %.0f seconds to allow Observers to print.\n", sleepDuration.Seconds())
	// time.Sleep(sleepDuration)
	// fmt.Println("Done.")

}
