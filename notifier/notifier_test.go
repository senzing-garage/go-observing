package notifier

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing-garage/go-observing/subject"
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
	var err error = nil
	return err
}

func teardown() error {
	var err error = nil
	return err
}

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNotify(test *testing.T) {
	ctx := context.TODO()
	observers := &subject.SubjectImpl{}
	origin := "Machine: 6 Process: Rover"
	subjectId := 1
	messageId := 2
	var err error = nil
	details := map[string]string{
		"data": "aData",
		"time": "aTime",
	}
	Notify(ctx, observers, origin, subjectId, messageId, err, details)
}
