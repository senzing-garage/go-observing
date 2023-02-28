package subject

import (
	"context"

	"github.com/senzing/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Subject interface {
	GetObservers(ctx context.Context) []observer.Observer
	HasObservers(ctx context.Context) bool
	NotifyObservers(ctx context.Context, message string) error
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}
