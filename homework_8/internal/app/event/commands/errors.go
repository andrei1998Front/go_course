package commands

import "errors"

var (
	ErrInvalidUUID  = errors.New("invalid event ID")
	ErrInvalidTitle = errors.New("invalid event title")
	ErrInvalidDate  = errors.New("invalid event date")
	ErrEmptyQuery   = errors.New("query is empty")
)
