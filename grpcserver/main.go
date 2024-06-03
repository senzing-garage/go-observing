package grpcserver

import (
	"context"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

type GrpcServer interface {
	GracefulStop(ctx context.Context) error
	Serve(ctx context.Context) error
}

// ----------------------------------------------------------------------------
// Constants
// ----------------------------------------------------------------------------

// Identfier of the  package found messages having the format "senzing-6464xxxx".
const ComponentID = 6464
