package main

import "fmt"

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func rich() {
	// EXERCISE: Richter Scale
	//
	var magnitude float64
	fmt.Println("Enter Richter Scale: ")
	_, err := fmt.Scanln(&magnitude)

	switch {
	case err != nil:
		fmt.Println("Error: ", err)

	case magnitude < 2.0:
		fmt.Println(magnitude, " is micro")

	case magnitude == 2.0 && magnitude < 3.0:
		fmt.Println(magnitude, " is very minor")

	case magnitude == 3.0 && magnitude < 4.0:
		fmt.Println(magnitude, " is minor")
	case magnitude == 4.0 && magnitude < 5.0:
		fmt.Println(magnitude, " is light")
	case magnitude == 5.0 && magnitude < 6.0:
		fmt.Println(magnitude, " is moderate")
	case magnitude == 6.0 && magnitude < 7.0:
		fmt.Println(magnitude, " is strong")
	case magnitude == 7.0 && magnitude < 8.0:
		fmt.Println(magnitude, " is major")
	case magnitude == 8.0 && magnitude < 10.0:
		fmt.Println(magnitude, " is great")
	case magnitude >= 10.0:
		fmt.Println(magnitude, " is massive")
	}
}
