package observer_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-observing/observer"
	"github.com/senzing-garage/go-observing/observerpb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const grpcAddress = "localhost:8261"

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestObserverGrpc_GetObserverId(test *testing.T) {
	ctx := context.TODO()
	grpcConnection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(test, err)
	grpcClient := observerpb.NewObserverClient(grpcConnection)
	observer := &observer.GrpcObserver{
		GrpcClient: grpcClient,
		ID:         "1",
	}
	assert.Equal(test, "1", observer.GetObserverID(ctx))
}

func TestObserverGrpc_UpdateObserver(test *testing.T) {
	_ = test
	ctx := context.TODO()
	grpcConnection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(test, err)
	grpcClient := observerpb.NewObserverClient(grpcConnection)
	observer := &observer.GrpcObserver{
		GrpcClient: grpcClient,
		ID:         "1",
	}
	observer.UpdateObserver(ctx, "A message")
}
