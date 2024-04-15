package commands

import "errors"

var (
	ErrInvalidUUID  = errors.New("invalid event ID")
	ErrInvalidTitle = errors.New("invalid event title")
)
