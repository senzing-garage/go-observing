package subject

import (
	"context"
	"sync"

	"github.com/senzing-garage/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SimpleSubject is the default implementation of the Subject interface.
type SimpleSubject struct {
	observerList []observer.Observer
	Lock         sync.RWMutex
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The GetObservers method returns a clone of registered Observers.
The caller will not have access to the type-struct's internal slice of observers.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *SimpleSubject) GetObservers(ctx context.Context) []observer.Observer {
	_ = ctx
	result := make([]observer.Observer, len(subject.observerList))
	copy(result, subject.observerList)

	return result
}

/*
The HasObservers method is used to determine if there are any Observers.
It can be used as a "guard" to avoid creating a message only to have
nothing observe it.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *SimpleSubject) HasObservers(ctx context.Context) bool {
	_ = ctx

	return len(subject.observerList) > 0
}

/*
The NotifyObservers method notifies all listeners of the message.
This is done asynchronously using goroutines.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (subject *SimpleSubject) NotifyObservers(ctx context.Context, message string) error {
	var err error

	for _, observer := range subject.observerList {
		go observer.UpdateObserver(ctx, message)
	}

	return err
}

/*
The RegisterObserver method adds an observer.Observer to an internal list of
listeners to be notified. The listener will receive event notifications.

Input
  - ctx: A context to control lifecycle.
  - observer: A component wanting to listen to events.
*/
func (subject *SimpleSubject) RegisterObserver(
	ctx context.Context,
	observer observer.Observer,
) error {
	var err error

	if observer != nil {
		subject.Lock.RLock()
		defer subject.Lock.RUnlock()
		if !contains(ctx, subject.observerList, observer) {
			subject.observerList = append(subject.observerList, observer)
		}
	}

	return err
}

/*
The UnregisterObserver method removes an observer.Observer from an internal list of
listeners.  The listener will no longer receive event notifications.

Input
  - ctx: A context to control lifecycle.
  - observer: A component no longer wanting to listen to events.
*/
func (subject *SimpleSubject) UnregisterObserver(
	ctx context.Context,
	observer observer.Observer,
) error {
	var err error

	if observer != nil {
		subject.Lock.RLock()
		defer subject.Lock.RUnlock()
		subject.observerList = removeFromSlice(ctx, subject.observerList, observer)
	}

	return err
}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func contains(ctx context.Context, haystack []observer.Observer, needle observer.Observer) bool {
	_ = ctx

	needleID := needle.GetObserverID(ctx)
	for _, value := range haystack {
		if value.GetObserverID(ctx) == needleID {
			return true
		}
	}

	return false
}

func removeFromSlice(
	ctx context.Context,
	observerList []observer.Observer,
	observerToRemove observer.Observer,
) []observer.Observer {
	removeID := observerToRemove.GetObserverID(ctx)
	observerListLength := len(observerList)

	for i, observer := range observerList {
		if observer.GetObserverID(ctx) == removeID {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]

			return observerList[:observerListLength-1]
		}
	}

	return observerList
}
