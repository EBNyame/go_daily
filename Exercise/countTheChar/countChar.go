package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "İNANÇ"

	length := utf8.RuneCountInString(name)

	fmt.Println(length)
}
