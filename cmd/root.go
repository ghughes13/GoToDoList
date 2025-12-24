package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "~~ Welcome To Your ToDo List ~~",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	for {
		showTodos("todo.txt")

		fmt.Println("What would you like to do? \n 1) Add \n 2) delete \n 3) exit")
		fmt.Println(" ")
		fmt.Print("Enter choice: ")

		text := getUserInput()

		if text == "1" {
			addToDo()
		} else if text == "2" {
			deleteToDo()
		} else if text == "3" {
			fmt.Println("Exiting...")
			os.Exit(0)
		} else {
			fmt.Println(" ")
			fmt.Println("!!!! Invalid choice. Please enter 1, 2, or 3.!!!!")
			fmt.Println(" ")
		}
	}
}
