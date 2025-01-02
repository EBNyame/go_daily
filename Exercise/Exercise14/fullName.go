package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func fullName() {
	// BONUS: Use a variable for the format specifier
	var name string
	fmt.Println("Name: ")
	fmt.Scan(&name)

	var lastname string
	fmt.Println("Lastname: ")
	fmt.Scan(&lastname)

	fmt.Printf("Your name is %s and your lastname is %s\n", name, lastname)
}
