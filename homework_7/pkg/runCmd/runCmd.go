package runcmd

import "fmt"

func RunCmd(cmd []string, env map[string]string) int {
	fmt.Println(env)

	return 0
}
