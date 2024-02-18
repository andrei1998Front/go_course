/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	rd "github.com/andrei1998Front/go_course/homework_7/pkg/readDir"
	rc "github.com/andrei1998Front/go_course/homework_7/pkg/runCmd"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "homework_7",
	Short: "Запуск программы с утановкой переменных окружения из файла",
	Long:  "Эта утилита позволяет запускать программы получая переменные окружения из определенной директории. См man envdir \nПример homework_7 /path/to/evndir command arg1 arg2",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println(cmd.Long)
			return
		}
		env, err := rd.ReadDir(args[0])

		if err != nil {
			log.Fatal(err)
		}

		i, err := rc.RunCmd(args[1:], env)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Код выхода программы: ", i)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
