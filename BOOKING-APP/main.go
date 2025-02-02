package main

import "fmt"

func main() {
	const totalTicket = 50
	var remainingTickets = 50
	var conferenceName = "Go Conference"
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", totalTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		var firstName string
		fmt.Println("first name: ")
		fmt.Scan(&firstName)

		var lastName string
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		var userMail string
		fmt.Println("Enter your email: ")
		fmt.Scan(&userMail)

		var userTickets int
		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		var totalRemainingTickets int
		remainingTickets = totalTicket - userTickets
		fmt.Println(remainingTickets)
		totalRemainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("%v remaining tickets for the %v \n", totalRemainingTickets, conferenceName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, userMail)

		fmt.Printf("This is all our bookings: %v \n", bookings)
	}
}
