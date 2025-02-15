package main

import "fmt"

// Your name
// Your favorite programming language
// "I'm learning Go for backend development!"

func greetUser(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	fmt.Println("My name is Exodus Blessed Nyame")
	fmt.Println("I'm learning Go for backend development!")

	// variable()
	greetUser("Exodus")

	results := multiply(5, 2)
	myName := fullName("Exodus", "Nyame")

	fmt.Printf("my name is %v and the results my final result is %v. \n", myName, results)
}

func multiply(x int, y int) int {
	return x * y
}

func fullName(firstNmane string, lastName string) string {
	return firstNmane + " " + lastName
}
