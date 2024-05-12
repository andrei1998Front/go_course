package config

import "errors"

var (
	ErrorInvalidNumberArgs = errors.New("invalid number of arguments")
	ErrorInvalidPortValue  = errors.New("port value must be an integer")
	ErrorOutPortRange      = errors.New("the port value must be in the range from 0 to 65535")
	ErrorInvalidTimeout    = errors.New("invalid timeout value")
)
