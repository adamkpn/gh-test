package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Check if there are command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a string to convert to uppercase.")
		return
	}

	// Join all arguments into a single string (in case of multiple words)
	input := strings.Join(os.Args[1:], " ")

	// Convert the input to uppercase
	uppercaseInput := strings.ToUpper(input)

	// Print the uppercase string
	fmt.Println(uppercaseInput)
}
