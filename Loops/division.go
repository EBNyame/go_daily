package main

import "fmt"

func division() {
	var number int
	fmt.Println("Enter the number")
	fmt.Scanln(&number)
	fmt.Println()

	fmt.Printf("%5s", "÷")
	for i := 0; i <= number; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= number; i++ {
		fmt.Printf("%5d", i)
		for j := 0; j <= number; j++ {
			fmt.Printf("%5d", i/j)
		}
		fmt.Println()

	}
}
