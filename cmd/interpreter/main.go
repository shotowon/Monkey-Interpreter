package main

import (
	"fmt"
	"log"
	"monkey/internal/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hi %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Enter \"quit\" to exit program.\n")
	repl.Start(os.Stdin, os.Stdout)
}
