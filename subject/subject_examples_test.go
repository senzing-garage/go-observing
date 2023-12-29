//go:build linux

package subject

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

//func ExampleSubjectImpl_HasObservers() {
//	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
//	ctx := context.TODO()
//	subject := &SubjectImpl{}
//	fmt.Print(subject.HasObservers(ctx))
//	// Output: false
//}

func ExampleSubjectImpl_RegisterObserver() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := &SubjectImpl{}
	observer := &observer.ObserverNull{
		Id: "Observer 1",
	}
	err := subject.RegisterObserver(ctx, observer)
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}

func ExampleSubjectImpl_NotifyObservers() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := &SubjectImpl{}
	err := subject.NotifyObservers(ctx, "Message 1")
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}

func ExampleSubjectImpl_UnregisterObserver() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := &SubjectImpl{}
	observer := &observer.ObserverNull{
		Id: "Observer 1",
	}
	err := subject.UnregisterObserver(ctx, observer)
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}
