package config

import (
	"fmt"
	"strconv"
	"time"

	flag "github.com/spf13/pflag"
)

type Config struct {
	Timeout time.Duration
	Port    int
}

//go:generate go run github.com/vektra/mockery/v2@v2.42.1 --name=FlagGetter
type FlagGetter interface {
	String(name string, value string, usage string) *string
	Args() []string
}

func New() *Config {
	return &Config{}
}

func (cfg *Config) Init() error {
	timeOutString := flag.String("timeout", "5s", "Timeout for server")
	flag.Parse()
	argsList := flag.Args()

	timeOut, err := setTimeOut(*timeOutString)
	if err != nil {
		return err
	}

	port, err := setPort(argsList)
	if err != nil {
		return err
	}

	cfg.Port = port
	cfg.Timeout = timeOut
	return nil
}

func setTimeOut(toString string) (time.Duration, error) {
	timeOut, err := time.ParseDuration(toString)

	if err != nil {
		return timeOut, fmt.Errorf("%w: %w", ErrorInvalidTimeout, err)
	}

	return timeOut, err
}

func setPort(args []string) (int, error) {

	if len(args) > 1 {
		return 0, ErrorManyArgs
	}

	if len(args) == 0 {
		return 8080, nil
	}

	port, err := strconv.Atoi(args[0])

	if err != nil {
		return 0, fmt.Errorf("%w: %w", ErrorInvalidPortValue, err)
	}

	if port < 0 || port > 65535 {
		return 0, ErrorOutPortRange
	}

	return port, nil
}
