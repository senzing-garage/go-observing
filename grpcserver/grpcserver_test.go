package grpcserver

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/senzing-garage/go-observing/observerpb"
	"github.com/senzing-garage/go-observing/subject"
	"github.com/stretchr/testify/require"
)

// ----------------------------------------------------------------------------
// Test harness
// ----------------------------------------------------------------------------

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	code := m.Run()
	err = teardown()
	if err != nil {
		fmt.Print(err)
	}
	os.Exit(code)
}

func setup() error {
	var err error
	return err
}

func teardown() error {
	var err error
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestSimpleGrpcServer_Serve(test *testing.T) {
	ctx := context.TODO()
	aSubject := &subject.SimpleSubject{}
	aGrpcServer := &SimpleGrpcServer{
		Port:    8260,
		Subject: aSubject,
	}

	go func() {
		err := aGrpcServer.Serve(ctx)
		require.NoError(test, err)
	}()
	time.Sleep(1 * time.Second)

	err := aGrpcServer.GracefulStop(ctx)
	require.NoError(test, err)
}

func TestSimpleGrpcServer_UpdateObserver(test *testing.T) {
	ctx := context.TODO()
	aSubject := &subject.SimpleSubject{}
	aGrpcServer := &SimpleGrpcServer{
		Port:    8260,
		Subject: aSubject,
	}
	request := &observerpb.UpdateObserverRequest{}
	_, err := aGrpcServer.UpdateObserver(ctx, request)
	require.NoError(test, err)
}
