package observer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

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
