package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// 	const GameUsage = `Welcome to the Lucky Number Game! 🍀

	// The program will pick %d random numbers.
	// Your mission is to guess one of those numbers.

	// The greater your number is, harder it gets.

	// Wanna play?`

	// 	const maxTurns = 5

	// 	fmt.Println(GameUsage)
	var answer string
	fmt.Println("yes/no")
	_, err := fmt.Scan(&answer)

	if err != nil {
		fmt.Println("wrong option Enter yes or no")
		return
	}

	con_answer := strings.ToLower(answer)

	if con_answer == "no" {
		return
	} else {
		var pick int
		fmt.Println("Pick a number: ")
		_, err := fmt.Scan(&pick)

		if err != nil {
			fmt.Println("Invalid input!")
		}

		guess := rand.Intn(pick + 1)

		if pick == guess {
			fmt.Println("You're lucky")
			return
		} else {
			fmt.Printf("Try again! the hidden number is: %d \n", guess)
		}
	}

}
