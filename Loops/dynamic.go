package main

import "fmt"

func dynamic() {
	var number int
	fmt.Println("Enter the size of table")
	fmt.Scanln(&number)
	fmt.Printf("you inputed %d multiplication table", number)
	fmt.Println()

	fmt.Printf("%5s", "X")

	for i := 0; i <= number; i++ {
		fmt.Printf("%5d", i)
		// fmt.Println()
	}
	fmt.Println()

	for i := 0; i <= number; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= number; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
