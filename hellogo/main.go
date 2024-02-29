package main

import (
	"fmt"

	"github.com/skodkani/Go-project/mystrings"
)

// Main function - the entry point to the Application.
func main() {
	// myString.Reverse method reverse the input string and printed on the STDOUT.
	fmt.Println(
		mystrings.Reverse("Hello World"),
	)
}
