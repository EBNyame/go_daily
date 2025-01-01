package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// ---------------------------------------------------------
// EXERCISE: Improved Banger
//
//  Change the Banger program the work with Unicode
//  characters.
//
// INPUT
//  "İNANÇ"
//
// EXPECTED OUTPUT
//  İNANÇ!!!!!
// ---------------------------------------------------------

func main() {
	var msg string
	fmt.Println("your name: ")
	fmt.Scan(&msg)

	s := msg + strings.Repeat("!", utf8.RuneCountInString(msg))

	fmt.Println(s)
}
