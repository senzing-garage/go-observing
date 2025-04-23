package notifier_test

import (
	"testing"

	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNotify(test *testing.T) {
	test.Parallel()
	ctx := test.Context()
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
}
