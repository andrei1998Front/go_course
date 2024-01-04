package scanner

import (
	"bufio"
	"os"
)

func ScanUserInput() string {
	var userInput string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		userInput = scanner.Text()
	}

	return userInput
}
