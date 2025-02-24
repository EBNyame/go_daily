package main

import "fmt"

func makeMap() {
	ages := make(map[string]int)
	ages["Exo"] = 12
	ages["kood"] = 33

	fmt.Println("Ages: ", ages)
	fmt.Println("Age of Exo:", ages["kofi"])

}
