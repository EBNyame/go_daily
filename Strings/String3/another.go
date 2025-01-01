package main

import "fmt"

func another() {
	const (
		EST = -(5 + iota)

		_
		MST
		PST
	)
	fmt.Println("This is another function")
	fmt.Println(EST, MST, PST)
}
