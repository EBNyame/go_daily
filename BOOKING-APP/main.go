package main

import "fmt"

func main() {
	const totalTicket = 50
	var remainingTickets = 50
	var conferenceName = "Go Conference"
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", totalTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
