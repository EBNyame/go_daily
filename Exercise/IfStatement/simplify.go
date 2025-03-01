package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

func simplify() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	if radius >= 200 && isSphere {
		fmt.Println("It's a big sphere")
	} else {
		fmt.Println("I don't know.")
	}
}
