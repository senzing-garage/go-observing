package observer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
