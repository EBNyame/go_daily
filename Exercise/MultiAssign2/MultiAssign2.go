package main

import "fmt"

func main() {

	var (
		planet string
		isTrue bool
		temp   float64
	)
	planet = "Mars"
	temp = 19.5
	isTrue = true

	fmt.Println("Air is good on", planet)
	fmt.Println("It's", isTrue)
	fmt.Println("It is", temp)
}
