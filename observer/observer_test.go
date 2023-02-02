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
