package main

import (
	"fmt"
	"os"
	"strings"
)

const usage = '[command] [string]
Available commands: lower, upper and title'



// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func SM() {
	args := os.Args
	if len(args) != 3{
		fmt.Println(usage)
        return
	}

	cmd, str := args[1], args[2]
	switch cmd {
		case "lower":
			fmt.Println(strings.ToLower(str))
		case "upper":
			fmt.Println(strings.ToUpper(str))
		case "title":
			fmt.Println(strings.Title(str))
		default:
			fmt.Println("Unknown command:", cmd)
	}
}