// In Go, code executed as an application must be in a main package.
package main

// import packages
import (
	"fmt"

	"github.com/EwanEwart/greetings"
	// "golang.org/x/text/message"
)

func main() {
	// Get a greeting by calling the greetings package’s Hello function.
	message := greetings.Hello("Kate")
	fmt.Println(message)
}
