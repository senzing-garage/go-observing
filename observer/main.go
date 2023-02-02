package observer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Observer interface {
	GetObserverId(ctx context.Context) string
	UpdateObserver(ctx context.Context, message string)
}
