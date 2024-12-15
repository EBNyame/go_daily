package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	name := "Exodus"
	kName := "맥저지"

	fmt.Println(utf8.RuneCountInString(kName))
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println(kName, ":", len(kName))

}
