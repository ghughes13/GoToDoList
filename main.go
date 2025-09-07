package main

import (
	"fmt"

	"github.com/ghughes13/GoToDoList/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("CLI executed")
}