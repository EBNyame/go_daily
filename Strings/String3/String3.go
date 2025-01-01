package main

import "fmt"

func main() {
	const (
		day = iota
		monday
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)
	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	another()
}
