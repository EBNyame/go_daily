package main

import "fmt"

// Create a slice of integers and print only the even numbers.
func assignment1() {
	integer := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for integer1 := range integer {
		if integer1%2 == 1 {
			fmt.Print(integer1)
		}
	}
}
