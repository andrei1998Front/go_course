package scanner

import (
	"bufio"
	"os"
)

func ScanUserInput() (string, error) {
	scanner := bufio.NewReader(os.Stdin)

	userInput, err := scanner.ReadString('\n')

	if err != nil {
		return "", err
	}

	return userInput, nil
}
