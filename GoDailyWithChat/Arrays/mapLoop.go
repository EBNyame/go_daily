package main

import "fmt"

func mapLoop() {
	user := map[string]string{
		"Name":  "Exodus",
		"Role":  "Backend Developer",
		"Stack": "Golang with Gin",
	}
	for key, value := range user {
		fmt.Println(key, ":", value)
	}
}
