package notifier

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/senzing-garage/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
Notify registered observers with an opinionated message in the form:

	{
		"origin": "user-defined identification of message origin",
		"subjectId": int,
		"messageId": int,
		"messageTime:"; int64,
		:
	}

Where the ":" represents key/values which is message specific.

Input

  - ctx: A context to control lifecycle.

  - observers: The observers to notify.

  - origin: Identifies the originator of the observer message. Empty string is OK.

  - subjectId: Identifies the component sending the message. See https://github.com/senzing-garage/knowledge-base/blob/main/lists/senzing-product-ids.md

  - messageId: The unique message sent by the component

  - err: Either nil for no error or an error object

  - details:  A map of key/value pairs to add to the notification message

Output
  - A string containing the error received from Senzing's G2Product.
*/
func Notify(
	ctx context.Context,
	subject subject.Subject,
	origin string,
	subjectID int,
	messageID int,
	err error,
	details map[string]string,
) {
	if subject == nil {
		return
	}

	if len(origin) > 0 {
		details["origin"] = origin
	}

	details["subjectId"] = strconv.Itoa(subjectID)
	details["messageId"] = strconv.Itoa(messageID)
	details["messageTime"] = time.Now().UTC().Format(time.RFC3339Nano)

	if err != nil {
		details["error"] = err.Error()
	}

	message, err := json.Marshal(details)
	if err != nil {
		panic(err)
	} else if subject != nil { // For safety.
		err := subject.NotifyObservers(ctx, string(message))
		if err != nil {
			panic(err)
		}
	}
}
