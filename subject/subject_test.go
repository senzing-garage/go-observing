package subject

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	subjectSingleton Subject
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) Subject {
	_ = ctx
	_ = test
	if subjectSingleton == nil {
		subjectSingleton = &SimpleSubject{}
	}
	return subjectSingleton
}

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
	var err error
	return err
}

func teardown() error {
	var err error
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSubjectImpl_GetObservers(test *testing.T) {
	ctx := context.TODO()
	subject := &SimpleSubject{}
	observer := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := subject.RegisterObserver(ctx, observer)
	require.NoError(test, err)
	err = subject.RegisterObserver(ctx, observer)
	require.NoError(test, err)
	err = subject.RegisterObserver(ctx, observer)
	require.NoError(test, err)

	observers := subject.GetObservers(ctx)
	assert.Len(test, observers, 1)
}

func TestSubjectImpl_GetObservers_multi(test *testing.T) {
	ctx := context.TODO()
	subject := &SimpleSubject{}
	observer1 := &observer.NullObserver{
		ID: "Observer 1",
	}
	observer2 := &observer.NullObserver{
		ID: "Observer 2",
	}
	err := subject.RegisterObserver(ctx, observer1)
	require.NoError(test, err)
	err = subject.RegisterObserver(ctx, observer1)
	require.NoError(test, err)
	err = subject.RegisterObserver(ctx, observer2)
	require.NoError(test, err)
	err = subject.RegisterObserver(ctx, observer2)
	require.NoError(test, err)

	observers := subject.GetObservers(ctx)
	assert.Len(test, observers, 2)
}

func TestSubjectImpl_HasObservers(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	assert.False(test, subject.HasObservers(ctx))
}

func TestSubjectImpl_RegisterObserver(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	observer := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := subject.RegisterObserver(ctx, observer)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))
}

func TestSubjectImpl_NotifyObservers(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	err := subject.NotifyObservers(ctx, "Message 1")
	require.NoError(test, err)
}

func TestSubjectImpl_UnregisterObserver(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	observer := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := subject.UnregisterObserver(ctx, observer)
	require.NoError(test, err)
	assert.False(test, subject.HasObservers(ctx))
}

func TestSubjectImpl_PrintBuffer(test *testing.T) {
	// This is a work-around for testing on windows.
	// This clears out the print buffer.
	// If not cleared, ExampleSubjectImpl_HasObservers() prints the buffer and confuses the "Output:"
	_ = test
	fmt.Print("")
}
