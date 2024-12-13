package main

import "fmt"

func main() {
	fmt.Println("What's your name??")

	var myName string

	fmt.Scan(&myName)

	fmt.Println("my name is ", myName)
}
