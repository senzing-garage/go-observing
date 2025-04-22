//go:build linux

package notifier_test

import (
	"context"

	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Examples for godoc documentation
// ----------------------------------------------------------------------------

func ExampleNotify() {
	// For more information, visit https://github.com/senzing-garage/go-observing/blob/main/notifier/notifier_examples_test.go
	ctx := context.TODO()
	subject := subject.NewSimpleSubject()
	origin := "Machine: 6 Process: Rover"
	subjectID := 1
	messageID := 2
	var err error
	details := map[string]string{
		"data": "aData",
		"time": "aTime",
	}
	notifier.Notify(ctx, subject, origin, subjectID, messageID, err, details)
	// Output:
}
