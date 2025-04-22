package notifier_test

import (
	"context"
	"testing"

	"github.com/senzing-garage/go-observing/notifier"
	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Test interface functions
// ----------------------------------------------------------------------------

func TestNotify(test *testing.T) {
	_ = test
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
}
