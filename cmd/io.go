package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func showTodos(filename string) {
	fmt.Println("------------------------------")
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No todos found!")
			return
		}
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	i := 1
	for scanner.Scan() {
		fmt.Printf("%d. %s\n", i, scanner.Text())
		i++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("------------------------------")
	fmt.Println(" ")
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return strings.TrimRight(text, "\r\n\n")
}
