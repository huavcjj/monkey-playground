package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/huavcjj/monkey-playground/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! Welcome to the Monkey programming language!\n", user.Username)
	fmt.Printf("Type 'exit' to quit.\n")

	repl.Start(os.Stdin, os.Stdout)
}
