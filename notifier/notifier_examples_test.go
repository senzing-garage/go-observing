//go:build linux

package notifier

import (
	"context"

	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNotify() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/notifier/notifier_examples_test.go
	ctx := context.TODO()
	observers := &subject.SimpleSubject{}
	origin := "Machine: 6 Process: Rover"
	subjectID := 1
	messageID := 2
	var err error
	details := map[string]string{
		"data": "aData",
		"time": "aTime",
	}
	Notify(ctx, observers, origin, subjectID, messageID, err, details)
	// Output:
}
