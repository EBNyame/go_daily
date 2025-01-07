package main

import (
	"fmt"
	"strconv"
)

func error1() {
	s := strconv.Itoa(42)
	fmt.Println(s)
	fmt.Printf("%T \n", s)
}
