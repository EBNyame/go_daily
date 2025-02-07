package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func TwoSums() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum: ", sum)
}
