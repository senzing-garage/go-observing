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

func TestObserverNull_GetObserverId(test *testing.T) {
	ctx := context.TODO()
	observer := &NullObserver{
		ID: "1",
	}
	assert.Equal(test, "1", observer.GetObserverID(ctx))
}

func TestObserverNull_UpdateObserver(test *testing.T) {
	_ = test
	ctx := context.TODO()
	observer := &NullObserver{
		ID: "1",
	}
	observer.UpdateObserver(ctx, "A message")
}

func TestObserverWhiteList_GetObserverId(test *testing.T) {
	ctx := context.TODO()
	observer := &WhiteListObserver{
		ID: "1",
	}
	assert.Equal(test, "1", observer.GetObserverID(ctx))
}

func TestObserverWhiteList_UpdateObserver(test *testing.T) {
	_ = test
	ctx := context.TODO()
	message11 := `{"subjectId":"1", "messageId": "1"}`
	message12 := `{"subjectId":"1", "messageId": "2"}`
	message21 := `{"subjectId":"2", "messageId": "1"}`
	message22 := `{"subjectId":"2", "messageId": "2"}`
	message31 := `{"subjectId":"3", "messageId": "1"}`

	observer := &WhiteListObserver{
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
}
