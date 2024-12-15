package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("What's your mood now?")
	var mood string
	fmt.Scan(&mood)

	fmt.Println("I'm feeling", mood+strings.Repeat("!", len(mood)))
}
