package grpcsubject

import (
	"context"

	"github.com/senzing/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type GrpcSubject interface {
	GetObservers(ctx context.Context) []observer.Observer
	HasObservers(ctx context.Context) bool
	NotifyObservers(ctx context.Context, message string) error
	RegisterObserver(ctx context.Context, observer observer.Observer) error
	Serve(ctx context.Context) error
	UnregisterObserver(ctx context.Context, observer observer.Observer) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6464xxxx".
const ProductId = 6464
