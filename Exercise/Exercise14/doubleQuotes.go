package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Double Quotes
//
//  Print "hello world" with double-quotes using Printf
//  (As you see here)
//
// NOTE
//  Output "shouldn't" be just: hello world
//
// EXPECTED OUTPUT
//  "hello world"
// ---------------------------------------------------------

func double() {
	fmt.Printf("%q\n", "hello world")
}
