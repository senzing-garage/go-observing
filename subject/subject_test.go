package subject

import (
	"context"
	"fmt"
	"testing"

	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	subjectSingleton Subject
)

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

func TestSubjectImpl_UnregisterObserver_Multiple(test *testing.T) {
	ctx := context.TODO()
	subject := &SimpleSubject{}
	assert.False(test, subject.HasObservers(ctx))

	observer1 := &observer.NullObserver{
		ID: "Observer 1",
	}
	err := subject.RegisterObserver(ctx, observer1)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))

	observer2 := &observer.NullObserver{
		ID: "Observer 2",
	}
	err = subject.RegisterObserver(ctx, observer2)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))

	observer3 := &observer.NullObserver{
		ID: "Observer 3",
	}
	err = subject.RegisterObserver(ctx, observer3)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))

	err = subject.UnregisterObserver(ctx, observer1)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))

	err = subject.UnregisterObserver(ctx, observer2)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))

	err = subject.UnregisterObserver(ctx, observer2)
	require.NoError(test, err)
	assert.True(test, subject.HasObservers(ctx))

	err = subject.UnregisterObserver(ctx, observer3)
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
