package main

import "fmt"

func slice() {
	number := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice", number)
	number = append(number, 6, 7, 8, 9)
	fmt.Println("Updated slice: ", number)
}
