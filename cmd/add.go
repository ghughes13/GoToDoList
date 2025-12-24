package cmd

import (
	"fmt"
	"os"
)

func addToDo() {
	fmt.Print("Enter text: ")
	text := getUserInput()

	f, err := os.OpenFile("todo.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(text + "\n"); err != nil {
		fmt.Fprintln(os.Stderr, "Error wriging File:", err)
		return
	}

	fmt.Println("\n Your ToDo List has been updated!")
}
