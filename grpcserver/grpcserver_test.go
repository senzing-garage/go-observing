package grpcserver_test

import (
	"context"
	"testing"
	"time"

	"github.com/senzing-garage/go-observing/grpcserver"
	"github.com/senzing-garage/go-observing/observerpb"
	"github.com/senzing-garage/go-observing/subject"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSimpleGrpcServer_Serve(test *testing.T) {
	ctx := context.TODO()
	aSubject := subject.NewSimpleSubject()
	aGrpcServer := &grpcserver.SimpleGrpcServer{
		Port:    8260,
		Subject: aSubject,
	}

	go func() {
		err := aGrpcServer.Serve(ctx)
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(1 * time.Second)

	err := aGrpcServer.GracefulStop(ctx)
	require.NoError(test, err)
}

func TestSimpleGrpcServer_UpdateObserver(test *testing.T) {
	ctx := context.TODO()
	aSubject := subject.NewSimpleSubject()
	aGrpcServer := &grpcserver.SimpleGrpcServer{
		Port:    8260,
		Subject: aSubject,
	}
	request := &observerpb.UpdateObserverRequest{}
	_, err := aGrpcServer.UpdateObserver(ctx, request)
	require.NoError(test, err)
}
