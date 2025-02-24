package main

import "fmt"

func arrLoop() {
	names := []string{"kofi", "Ama", "Kwame"}

	for i, names := range names {
		fmt.Println(i, "Name: ", names)
	}
}
