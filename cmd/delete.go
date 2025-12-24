package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func deleteToDo() {
	fmt.Print("Enter item number to delete: ")
	input := getUserInput()
	idx, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || idx <= 0 {
		fmt.Println("Invalid item number.")
		return
	}

	f, err := os.Open("todo.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	scanner := bufio.NewScanner(f)
	var lines []string
	i := 1
	for scanner.Scan() {
		if i != idx {
			lines = append(lines, scanner.Text())
		}
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		f.Close()
		return
	}
	f.Close() // close before reopening for write

	f, err = os.OpenFile("todo.txt", os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Error opening file for writing:", err)
		return
	}
	defer f.Close()

	for _, line := range lines {
		if _, err := f.WriteString(line + "\n"); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	fmt.Println("Item deleted.")
}
