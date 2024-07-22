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
		request a greetings message
		Assign both of the Hello return values,
		including the error, to variables
		Change the Hello argument from a specific name
		to an empty string,
		so you can try out the error-handling code.
	*/
	message, err := greetings.Hello("Kate")
	// message, err := greetings.Hello("")
	/*
		If an error was returned,
		print it to the console and exit the program.
		Look for a non-nil error value. There's no sense continuing in this case.
	*/
	if err != nil {
		/*
			Use the functions in the standard library's "log package"
			to output error information.
			If you receive an error,
			you use "the log package's "Fatal" function"
			to print the error and stop the program.
		*/
		log.Fatal(err)
	}
	/*
		If no error was returned,
		print the returned message on the console.
		Get a greeting by calling the greetings package’s Hello function.
	*/
	fmt.Println(message)
}
