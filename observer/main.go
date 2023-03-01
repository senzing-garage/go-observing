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

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6462xxxx".
const ProductId = 6462
