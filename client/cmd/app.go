/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"client/api"
	"client/utils"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func NewAppCommand(quizApi *api.QuizApi) *cobra.Command {
	return &cobra.Command{
		Use:   "quiz",
		Short: "This command starts the simple test",
		Run: func(cmd *cobra.Command, args []string) {
			appShell(quizApi)
		},
	}
}

func Execute(quizApi *api.QuizApi) {
	appCommand := NewAppCommand(quizApi)
	err := appCommand.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func appShell(quizApi *api.QuizApi) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter your name: ")
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input)
	state := quizApi.CreateUser(name)
	switch state {
	case utils.STATUS_USER_CREATED:
		{
			fmt.Println("Welcome,", name)
			break
		}
	case utils.STATUS_USER_EXISTS:
		{
			fmt.Println("Welcome back!", name)
			break
		}
	default:
		{
			fmt.Println("Welcome,", name)
			break
		}
	}
	fmt.Println("To start the test, use command `s`. To exit the test, use command `\\q`")
	for {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "\\q":
			{
				fmt.Println("Exiting the test.")
				return
			}
		case "s":
			{
				quizCmd := NewQuizCommand(quizApi, name)
				quizCmd.Execute()
				return
			}
		default:
			{
				fmt.Println("Unknown command, please use `s` to start or `\\q` for exiting")
			}
		}
	}
}

func init() {
	//quizCmd.AddCommand(quizCmd)
}
