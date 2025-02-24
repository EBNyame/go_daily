package main

import "fmt"

// Create a map of students and their scores, then print students who scored above 80.
func assignment2() {
	student := make(map[string]int)
	student["Aba"] = 60
	student["kofi"] = 70
	student["Selina"] = 90

	for key, value := range student {
		// fmt.Println("Key:", key, "Value:", value)

		if value > 80 {
			fmt.Println(key)
		}
	}
}
