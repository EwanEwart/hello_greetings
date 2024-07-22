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
	including the "log entry prefix"
	and a "flag" to disable printing the time, source file, and line number.
	*/
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	/*
		Create a names variable as a slice type holding three names.

		In your hello_greetings/hello_greetings.go calling code,
		pass a slice of names,
		then print the contents of the names/messages map you get back.
	*/
	// A slice of names
	names := []string{"Mater", "Pater", "Filia", "Kitty"}

	/*
		Pass the names variable as the argument
		to the Hellos function

		Request greeting messages for the names.
	*/
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	/*
		If no error was returned,
		print the returned map of messages to the console.
	*/
	fmt.Println(messages)

}
