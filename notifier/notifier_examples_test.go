//go:build linux

package notifier

import (
	"context"

	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNotify() {
	// For more information, visit https://github.com/Senzing/go-observing/blob/main/notifier/notifier_examples_test.go
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
	// Output:
}
