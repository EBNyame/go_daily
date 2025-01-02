package main

import "fmt"

func typee() {
	var number int
	var word string
	var boolean bool

	fmt.Printf("Number:%T\n Word: %T\n Boolean: %T\n", number, word, boolean)
}
