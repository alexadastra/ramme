package system

import (
	"errors"
)

// ErrNotImplemented declares error for method that isn't implemented
var ErrNotImplemented = errors.New("this method is not implemented")

// Operator defines reload, maintenance and shutdown interface
type Operator interface {
	Reload() error
	Maintenance() error
	Shutdown(error) error

	Add(func() error, func(error))
	Run() error
}

// actor represents execution and interruption of a program entity
type actor struct {
	execute   func() error
	interrupt func(error)
}

// GroupOperator implements Operator interface
type GroupOperator struct {
	actors       []actor
	actorsErrors chan error
}

// NewGroupOperator is a GroupOperator constructor
func NewGroupOperator() *GroupOperator {
	return &GroupOperator{
		actors: make([]actor, 0),
	}
}

// Add adds new actor to handle
func (gop *GroupOperator) Add(execute func() error, interrupt func(error)) {
	gop.actors = append(gop.actors, actor{execute: execute, interrupt: interrupt})
}

// Run starts execution for all handlers and shuts everything down if some of them broke
func (gop *GroupOperator) Run() error {
	gop.actorsErrors = make(chan error, len(gop.actors))
	if len(gop.actors) == 0 {
		return nil
	}

	// Run each actor.
	for _, a := range gop.actors {
		go func(a actor) {
			gop.actorsErrors <- a.execute()
		}(a)
	}

	// Wait for the first actor to stop.
	err := <-gop.actorsErrors

	return gop.Shutdown(err)
}

// Reload operation implementation
func (gop *GroupOperator) Reload() error {
	return ErrNotImplemented
}

// Maintenance operation implementation
func (gop *GroupOperator) Maintenance() error {
	return ErrNotImplemented
}

// Shutdown operation implementation
func (gop *GroupOperator) Shutdown(err error) error {
	// Signal all actors to stop.
	for _, a := range gop.actors {
		a.interrupt(err)
	}

	// Wait for all actors to stop.
	for i := 1; i < cap(gop.actorsErrors); i++ {
		<-gop.actorsErrors
	}

	// Return the original error.
	return err
}
