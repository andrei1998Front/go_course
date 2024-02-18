/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"

	rd "github.com/andrei1998Front/go_course/homework_7/pkg/readDir"

	"github.com/andrei1998Front/go_course/homework_7/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("fff")
	rd.ReadDir(".")
}
