package observer

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestObserverNull_GetObserverId(test *testing.T) {
	ctx := context.TODO()
	observer := &ObserverNull{
		Id: "1",
	}
	assert.Equal(test, "1", observer.GetObserverId(ctx))
}

func TestObserverNull_UpdateObserver(test *testing.T) {
	ctx := context.TODO()
	observer := &ObserverNull{
		Id: "1",
	}
	observer.UpdateObserver(ctx, "A message")
}

func TestObserverWhiteList_GetObserverId(test *testing.T) {
	ctx := context.TODO()
	observer := &ObserverWhiteList{
		Id: "1",
	}
	assert.Equal(test, "1", observer.GetObserverId(ctx))
}

func TestObserverWhiteList_UpdateObserver(test *testing.T) {
	ctx := context.TODO()
	message11 := `{"subjectId":"1", "messageId": "1"}`
	message12 := `{"subjectId":"1", "messageId": "2"}`
	message21 := `{"subjectId":"2", "messageId": "1"}`
	message22 := `{"subjectId":"2", "messageId": "2"}`
	message31 := `{"subjectId":"3", "messageId": "1"}`

	observer := &ObserverWhiteList{
		Id: "1",
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
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleObserverNull_GetObserverId() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/observer/observer_test.go
	ctx := context.TODO()
	observer := &ObserverNull{
		Id: "1",
	}
	fmt.Print(observer.GetObserverId(ctx))
	// Output: 1
}

func ExampleObserverNull_UpdateObserver() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/observer/observer_test.go
	ctx := context.TODO()
	observer := &ObserverNull{
		Id: "1",
	}
	observer.UpdateObserver(ctx, "A message")
	// Output: Observer: 1;  Message: A message
}

func ExampleObserverWhiteList_GetObserverId() {
	ctx := context.TODO()
	observer := &ObserverWhiteList{
		Id: "1",
	}
	fmt.Print(observer.GetObserverId(ctx))
	// Output: 1
}

func ExampleObserverWhiteList_UpdateObserver() {
	ctx := context.TODO()
	message11 := `{"subjectId":"1", "messageId": "1"}`
	message12 := `{"subjectId":"1", "messageId": "2"}`
	message21 := `{"subjectId":"2", "messageId": "1"}`
	message22 := `{"subjectId":"2", "messageId": "2"}`
	message31 := `{"subjectId":"3", "messageId": "1"}`

	observer := &ObserverWhiteList{
		Id: "1",
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
