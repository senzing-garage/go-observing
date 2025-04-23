package observer_test

import (
	"testing"

	"github.com/senzing-garage/go-observing/observer"
	"github.com/stretchr/testify/assert"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestObserverRaw_GetObserverId(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	observer := &observer.RawObserver{
		ID: "1",
	}
	assert.Equal(test, "1", observer.GetObserverID(ctx))
}

func TestObserverRaw_UpdateObserver(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
	observer := &observer.RawObserver{
		ID: "1",
	}
	observer.UpdateObserver(ctx, "A message")
}
