package config

import (
	"fmt"
	"strconv"
	"time"

	flag "github.com/spf13/pflag"
)

type Config struct {
	Timeout time.Duration
	Host    string
	Port    int
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) Init() error {
	timeOutString := flag.String("timeout", "10s", "Timeout for server")
	flag.Parse()
	argsList := flag.Args()

	if len(argsList) != 2 {
		return ErrorInvalidNumberArgs
	}

	timeOut, err := setTimeOut(*timeOutString)
	if err != nil {
		return err
	}

	port, err := setPort(argsList[1])
	if err != nil {
		return err
	}

	cfg.Port = port
	cfg.Timeout = timeOut
	cfg.Host = argsList[0]
	return nil
}

func setTimeOut(toString string) (time.Duration, error) {
	timeOut, err := time.ParseDuration(toString)

	if err != nil {
		return timeOut, fmt.Errorf("%w: %w", ErrorInvalidTimeout, err)
	}

	return timeOut, err
}

func setPort(portValue string) (int, error) {
	port, err := strconv.Atoi(portValue)

	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrorInvalidPortValue, err)
	}

	if port < 0 || port > 65535 {
		return 0, ErrorOutPortRange
	}

	return port, nil
}
