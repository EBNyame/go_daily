package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func verbose() {
	var totalSum int
	for i := 1; i <= 10; i++ {
		totalSum += i
		fmt.Print(i)
		if i < 10 {
			fmt.Print(" + ")
		}

	}
	fmt.Printf(" = %d\n", totalSum)
}
