package main

import "fmt"

func error2() {
	var age int

	fmt.Println("Please enter a valid age")
	fmt.Scanln(&age)

	_, err := fmt.Scan(age)
	if err != nil {
		fmt.Println(err, "We got an error")
	} else {
		fmt.Println("You are", age, "years old")
	}
}
