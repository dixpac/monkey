package main

import (
	"fmt"
	"os"

	"dix/monkey/repl"
)

func main() {
	fmt.Printf("Hello from Monkey programming language!\n")
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
