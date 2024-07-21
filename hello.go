package main

//  In Go, code executed as an application must be in a main package.

import (
	"fmt"

	"github.com/EwanEwart/greetings"
	"golang.org/x/text/message"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Kate")
	fmt.Println(message)
}
