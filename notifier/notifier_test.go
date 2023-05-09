package notifier

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/senzing/go-observing/subject"
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
	origin := ""
	subjectId := 1
	messageId := 2
	var err error = nil
	details := map[string]string{
		"data": "aData",
		"time": "aTime",
	}
	Notify(ctx, observers, origin, subjectId, messageId, err, details)
}

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNotify() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/notifier/notifier_test.go
	ctx := context.TODO()
	observers := &subject.SubjectImpl{}
	origin := ""
	subjectId := 1
	messageId := 2
	var err error = nil
	details := map[string]string{
		"data": "aData",
		"time": "aTime",
	}
	Notify(ctx, observers, origin, subjectId, messageId, err, details)
	// Output:
}
