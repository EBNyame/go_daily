package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Raw Concat
//
//  1. Initialize the name variable
//     by getting input from the command line
//
//  2. Use concatenation operator to concatenate it
//     with the raw string literal below
//
// NOTE
//  You should concatenate the name variable in the correct
//  place.
//
// EXPECTED OUTPUT
//  Let's say that you run the program like this:
//   go run main.go inanç
//
//  Then it should output this:
//   hi inanç!
//   how are you?
// ---------------------------------------------------------

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	// replace and concatenate the `name` variable
	// after `hi ` below

	msg := `hi, ` + name +
		`  how are you?`

	fmt.Println(msg)
}
