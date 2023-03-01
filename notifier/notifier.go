package notifier

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/senzing/go-observing/subject"
)

// ----------------------------------------------------------------------------
// Public functions
// ----------------------------------------------------------------------------

/*
Notify registered observers with an opinionated message in the form:

	{
		"subjectId": int,
		"messageId": int,
		"messageTime:"; int64,
		:
	}

Where the ":" represents key/values which is message specific.

Input
  - ctx: A context to control lifecycle.
  - observers: The observers to notify.
  - productId: Identifies the component sending the message. See https://github.com/Senzing/knowledge-base/blob/main/lists/senzing-product-ids.md
  - messageId: The unique message sent by the component
  - err: Either nil for no error or an error object
  - details:  A map of key/value pairs to add to the notification message

Output
  - A string containing the error received from Senzing's G2Product.
*/
func Notify(ctx context.Context, observers subject.Subject, productId int, messageId int, err error, details map[string]string) {
	if observers != nil {
		now := time.Now()
		details["subjectId"] = strconv.Itoa(productId)
		details["messageId"] = strconv.Itoa(messageId)
		details["messageTime"] = strconv.FormatInt(now.UnixNano(), 10)
		if err != nil {
			details["error"] = err.Error()
		}
		message, err := json.Marshal(details)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			observers.NotifyObservers(ctx, string(message))
		}
	}
}
