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
	fmt.Println(strings.ToUpper(mood))

	fmt.Println(strings.Repeat("!", len(mood)), mood+strings.Repeat("!", len(mood)))
}
