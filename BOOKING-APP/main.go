package main

import "fmt"

func main() {
	const totalTicket = 50
	var remainingTickets = 50
	var conferenceName = "Go Conference"
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", totalTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings [50]string

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

	remainingTickets = totalTicket - userTickets
	bookings[0] = firstName + " " + lastName
	fmt.Println(bookings)

	fmt.Printf("%v remaining tickets for the %v \n", remainingTickets, conferenceName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, userMail)
}
