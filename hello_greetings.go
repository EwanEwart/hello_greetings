// In Go, code executed as an application must be in a main package.
package main

// import packages
import (
	"fmt"
	"log"

	"github.com/EwanEwart/greetings"
	// "golang.org/x/text/message"
)

func main() {
	/* Set properties of the predefined Logger,
	including the log entry prefix
	and a flag to disable printing the time, source file, and line number.
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greetings message
	message, err := greetings.Hello("")
	/*
		If an error was returned,
		print it to the console and exit the program.
	*/
	if err != nil {
		log.Fatal(err)
	}
	/*
		If no error was returned,
		print the returned message on the console.
		Get a greeting by calling the greetings package’s Hello function.
	*/
	fmt.Println(message)
}
