package observer

import (
	"context"
	"encoding/json"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// WhiteListObserver is a simple example of the Subject interface.
// It is mainly used for testing.
type WhiteListObserver struct {
	ID        string
	IsSilent  bool
	WhiteList map[int]map[int]bool
}

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

type MessageFormat struct {
	SubjectID string `json:"subjectId"`
	MessageID string `json:"messageId"`
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetObserverID method returns the unique identifier of the observer.
Use by the subject to manage the list of Observers.

Input
  - ctx: A context to control lifecycle.
*/
func (observer *WhiteListObserver) GetObserverID(ctx context.Context) string {
	_ = ctx
	return observer.ID
}

/*
The UpdateObserver method processes the message sent by the Subject.
The subject invokes UpdateObserver as a goroutine.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (observer *WhiteListObserver) UpdateObserver(ctx context.Context, message string) {
	_ = ctx
	if !observer.IsSilent {
		isOnWhiteList, err := observer.onWhiteList(message)
		if err != nil {
			outputf("Error: Observer: %s;  Message: %s; Error: %v\n", observer.ID, message, err)

		}
		if isOnWhiteList {
			outputf("Observer: %s;  Message: %s\n", observer.ID, message)
		}
	}
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (observer *WhiteListObserver) parseMessage(message string) (int, int, error) {
	var parsedMessage MessageFormat
	err := json.Unmarshal([]byte(message), &parsedMessage)
	if err != nil {
		return 0, 0, err
	}
	subjectID, err := strconv.Atoi(parsedMessage.SubjectID)
	if err != nil {
		return 0, 0, err
	}
	messageID, err := strconv.Atoi(parsedMessage.MessageID)
	if err != nil {
		return 0, 0, err
	}
	return subjectID, messageID, err
}

func (observer *WhiteListObserver) onWhiteList(message string) (bool, error) {
	var err error
	if !observer.IsSilent {
		subjectID, messageID, err := observer.parseMessage(message)
		if err != nil {
			return false, err
		}
		return observer.WhiteList[subjectID][messageID], err
	}
	return false, err
}
