package observer_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestObserverRaw_GetObserverId(test *testing.T) {
	ctx := context.TODO()
	observer := &observer.RawObserver{
		ID: "1",
	}
	assert.Equal(test, "1", observer.GetObserverID(ctx))
}

func TestObserverRaw_UpdateObserver(test *testing.T) {
	_ = test
	ctx := context.TODO()
	observer := &observer.RawObserver{
		ID: "1",
	}
	observer.UpdateObserver(ctx, "A message")
}
