package observer

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// ObserverNull is a simple example of the Subject interface.
// It is mainly used for testing.
type ObserverWhiteList struct {
	Id        string
	IsSilent  bool
	WhiteList map[int]map[int]bool
}

// ----------------------------------------------------------------------------
// Types - struct
// ----------------------------------------------------------------------------

type MessageFormat struct {
	SubjectId string `json:"subjectId"`
	MessageId string `json:"messageId"`
}

// ----------------------------------------------------------------------------
// Internal methods
// ----------------------------------------------------------------------------

func (observer *ObserverWhiteList) parseMessage(message string) (int, int, error) {
	var parsedMessage MessageFormat
	err := json.Unmarshal([]byte(message), &parsedMessage)
	if err != nil {
		return 0, 0, err
	}
	subjectId, err := strconv.Atoi(parsedMessage.SubjectId)
	if err != nil {
		return 0, 0, err
	}
	messageId, err := strconv.Atoi(parsedMessage.MessageId)
	if err != nil {
		return 0, 0, err
	}
	return subjectId, messageId, err
}

func (observer *ObserverWhiteList) onWhiteList(message string) (bool, error) {
	var err error = nil
	if !observer.IsSilent {
		subjectId, messageId, err := observer.parseMessage(message)
		if err != nil {
			return false, err
		}
		return observer.WhiteList[subjectId][messageId], err
	}
	return false, err
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetObserverId method returns the unique identifier of the observer.
Use by the subject to manage the list of Observers.

Input
  - ctx: A context to control lifecycle.
*/
func (observer *ObserverWhiteList) GetObserverId(ctx context.Context) string {
	return observer.Id
}

/*
The UpdateObserver method processes the message sent by the Subject.
The subject invokes UpdateObserver as a goroutine.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (observer *ObserverWhiteList) UpdateObserver(ctx context.Context, message string) {
	if !observer.IsSilent {
		isOnWhiteList, err := observer.onWhiteList(message)
		if err != nil {
			fmt.Printf("Error: Observer: %s;  Message: %s; Error: %v\n", observer.Id, message, err)

		}
		if isOnWhiteList {
			fmt.Printf("Observer: %s;  Message: %s\n", observer.Id, message)
		}
	}
}
