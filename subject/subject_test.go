package subject

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing/go-observing/observer"
	"github.com/stretchr/testify/assert"
)

var (
	subjectSingleton Subject
)

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func getTestObject(ctx context.Context, test *testing.T) Subject {
	if subjectSingleton == nil {
		subjectSingleton = &SubjectImpl{}
	}
	return subjectSingleton
}

func testError(test *testing.T, ctx context.Context, err error) {
	if err != nil {
		test.Log("Error:", err.Error())
		assert.FailNow(test, err.Error())
	}
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

func TestSubjectImpl_HasObservers(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	assert.False(test, subject.HasObservers(ctx))
}

func TestSubjectImpl_RegisterObserver(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	observer := &observer.ObserverNull{
		Id: "Observer 1",
	}
	err := subject.RegisterObserver(ctx, observer)
	testError(test, ctx, err)
	assert.True(test, subject.HasObservers(ctx))
}

func TestSubjectImpl_NotifyObservers(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	err := subject.NotifyObservers(ctx, "Message 1")
	testError(test, ctx, err)
}

func TestSubjectImpl_UnregisterObserver(test *testing.T) {
	ctx := context.TODO()
	subject := getTestObject(ctx, test)
	observer := &observer.ObserverNull{
		Id: "Observer 1",
	}
	err := subject.UnregisterObserver(ctx, observer)
	testError(test, ctx, err)
	assert.False(test, subject.HasObservers(ctx))
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleSubjectImpl_HasObservers() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := &SubjectImpl{}
	fmt.Print(subject.HasObservers(ctx))
	// Output: false
}

func ExampleSubjectImpl_RegisterObserver() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/subject/subject_test.go
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
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/subject/subject_test.go
	ctx := context.TODO()
	subject := &SubjectImpl{}
	err := subject.NotifyObservers(ctx, "Message 1")
	if err != nil {
		fmt.Print(err)
	}
	// Output:
}

func ExampleSubjectImpl_UnregisterObserver() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/subject/subject_test.go
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
