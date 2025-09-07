package main

import (
	"fmt"

	"github.com/ghughes13/gotodolist/cmd"
)

func main() {
	cmd.Execute()
	fmt.Println("CLI executed")
}