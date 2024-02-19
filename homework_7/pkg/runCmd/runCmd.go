package runcmd

import (
	"errors"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env map[string]string) (int, error) {
	lenCmd := len(cmd)

	if lenCmd == 0 {
		return 0, errors.New("ошибка. Команда не передана")
	}

	comand := exec.Command(cmd[0])

	if lenCmd > 1 {
		comand.Args = cmd[1:]
	}

	var envForCommand []string

	for key, value := range env {
		envForCommand = append(envForCommand, key+"="+value)
	}

	comand.Env = append(os.Environ(), envForCommand...)
	comand.Stdout = os.Stdout
	comand.Stderr = os.Stderr

	err := comand.Run()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.ExitCode(), errors.New("Ошибка выполнения команды: " + err.Error())
		}

		return 0, errors.New("Неизвестная команда: " + err.Error())
	}

	return 0, nil
}
