package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string = "Exodus"
	var myName string = "맥저지"
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println(utf8.RuneCountInString(myName))
	fmt.Println("wrong count of myname: ", len(myName))
}
