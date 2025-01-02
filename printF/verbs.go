package main

import "fmt"

func verbs() {
	var name string = "Exodus"
	var age int = 100
	var is_student bool = false

	fmt.Printf("your name is %[1]v you're %[2]v and is %[3]v that you're a student \n", name, age, is_student)
}
