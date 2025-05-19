//go:build linux

package observer_test

import (
	"context"
	"fmt"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNullObserver_GetObserverID() {
	// For more information,
	// visit https://github.com/senzing-garage/go-observing/blob/main/observer/observer_examples_test.go
	ctx := context.TODO()
	observer := &observer.NullObserver{
		ID: "1",
	}
	fmt.Print(observer.GetObserverID(ctx))
	// Output: 1
}

func ExampleNullObserver_UpdateObserver() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/observer/observer_test.go
	ctx := context.TODO()
	observer := &observer.NullObserver{
		ID: "1",
	}
	observer.UpdateObserver(ctx, "A message")
	// Output: Observer: 1;  Message: A message
}

func ExampleWhiteListObserver_GetObserverID() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/observer/observer_test.go
	ctx := context.TODO()
	observer := &observer.WhiteListObserver{
		ID: "1",
	}
	fmt.Print(observer.GetObserverID(ctx))
	// Output: 1
}

func ExampleWhiteListObserver_UpdateObserver() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/observer/observer_test.go
	ctx := context.TODO()
	message11 := `{"subjectId":"1", "messageId": "1"}`
	message12 := `{"subjectId":"1", "messageId": "2"}`
	message21 := `{"subjectId":"2", "messageId": "1"}`
	message22 := `{"subjectId":"2", "messageId": "2"}`
	message31 := `{"subjectId":"3", "messageId": "1"}`

	observer := &observer.WhiteListObserver{
		ID: "1",
		WhiteList: map[int]map[int]bool{
			1: {
				1: true,
				2: true,
			},
			2: {
				2: true,
			},
		},
	}
	observer.UpdateObserver(ctx, message11)
	observer.UpdateObserver(ctx, message12)
	observer.UpdateObserver(ctx, message21)
	observer.UpdateObserver(ctx, message22)
	observer.UpdateObserver(ctx, message31)
	// Output:
	// Observer: 1;  Message: {"subjectId":"1", "messageId": "1"}
	// Observer: 1;  Message: {"subjectId":"1", "messageId": "2"}
	// Observer: 1;  Message: {"subjectId":"2", "messageId": "2"}
}
