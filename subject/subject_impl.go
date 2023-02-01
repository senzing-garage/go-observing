package subject

import (
	"context"
	"sync"

	"github.com/senzing/go-observing/observer"
)

// ----------------------------------------------------------------------------
// Types
// ----------------------------------------------------------------------------

// SubjectImpl is the default implementation of the Subject interface.
type SubjectImpl struct {
	observerList []observer.Observer
}

// ----------------------------------------------------------------------------
// Variables
// ----------------------------------------------------------------------------

var lock = sync.RWMutex{}

// ----------------------------------------------------------------------------
// Internal functions
// ----------------------------------------------------------------------------

func removeFromSlice(ctx context.Context, observerList []observer.Observer, observerToRemove observer.Observer) []observer.Observer {
	removeId := observerToRemove.GetObserverId(ctx)
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observer.GetObserverId(ctx) == removeId {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

// ----------------------------------------------------------------------------
// Interface methods
// ----------------------------------------------------------------------------

/*
The HasObservers method is used to determine if there are any Observers.
It can be used as a "guard" to avoid creating a message only to have
nothing observe it.

Input
  - ctx: A context to control lifecycle.
*/
func (subject *SubjectImpl) HasObservers(ctx context.Context) bool {
	return len(subject.observerList) > 0
}

/*
The NotifyObservers method notifies all listeners of the message.
This is done asynchronously using goroutines.

Input
  - ctx: A context to control lifecycle.
  - message: The string to propagate to all registered Observers.
*/
func (subject *SubjectImpl) NotifyObservers(ctx context.Context, message string) error {
	var err error = nil
	lock.RLock()
	defer lock.RUnlock()
	for _, observer := range subject.observerList {
		observer.UpdateObserver(ctx, message)
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
func (subject *SubjectImpl) RegisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error = nil
	lock.RLock()
	defer lock.RUnlock()
	subject.observerList = append(subject.observerList, observer)
	return err
}

/*
The UnregisterObserver method removes an observer.Observer from an internal list of
listeners.  The listener will no longer receive event notifications.

Input
  - ctx: A context to control lifecycle.
  - observer: A component no longer wanting to listen to events.
*/
func (subject *SubjectImpl) UnregisterObserver(ctx context.Context, observer observer.Observer) error {
	var err error = nil
	lock.RLock()
	defer lock.RUnlock()
	subject.observerList = removeFromSlice(ctx, subject.observerList, observer)
	return err
}
