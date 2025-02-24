package main

import "fmt"

func sumArr() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, value := range arr {
		sum += value
	}
	fmt.Println("sum: ", sum)
}
