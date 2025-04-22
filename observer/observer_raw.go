package observer

import (
	"context"
	"fmt"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// RawObserver is a simple example of the Subject interface.
// It is mainly used for testing.
type RawObserver struct {
	ID       string
	IsSilent bool
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
func (observer *RawObserver) GetObserverID(ctx context.Context) string {
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
func (observer *RawObserver) UpdateObserver(ctx context.Context, message string) {
	_ = ctx
	if !observer.IsSilent {
		fmt.Println(message) //nolint
	}
}
