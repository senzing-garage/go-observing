package observer

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type Observer interface {
	GetObserverID(ctx context.Context) string
	UpdateObserver(ctx context.Context, message string)
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6462xxxx".
const ComponentID = 6462
