//go:build linux

package subject_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSimpleSubject_HasObservers() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := subject.NewSimpleSubject()
	fmt.Print(subject.HasObservers(ctx))
	// Output: false
}

func ExampleSimpleSubject_RegisterObserver() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := subject.NewSimpleSubject()
	observer := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := subject.RegisterObserver(ctx, observer)
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}

func ExampleSimpleSubject_NotifyObservers() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := subject.NewSimpleSubject()
	err := subject.NotifyObservers(ctx, "Message 1")
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}

func ExampleSimpleSubject_UnregisterObserver() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := subject.NewSimpleSubject()
	observer := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := subject.UnregisterObserver(ctx, observer)
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}
