package event

import "errors"

var (
	ErrNonExistentID    = errors.New("events with this ID do not exist")
	ErrNonExistentEvent = errors.New("no such event exists")
	ErrDateBusy         = errors.New("an event for this date has already been selected")
	ErrNonExistentDate  = errors.New("there is no event with this date")
	ErrExistentID       = errors.New("events with the same ID already exist")
)
