package main

import (
	"fmt"
	"strings"
)

func main() {
	const totalTicket = 50
	var remainingTickets = totalTicket
	var conferenceName = "Go Conference"

	greet(conferenceName, remainingTickets, totalTicket)

	var bookings []string

	for {
		firstName, lastName, userMail, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicket := validateUserInput(firstName, lastName, userMail, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {
			remainingTickets = remainingTickets - userTickets
			fmt.Println(remainingTickets)

			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("%v remaining tickets for the %v \n", remainingTickets, conferenceName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v \n", firstName, lastName, userTickets, userMail)

			fmt.Printf("This is all our bookings: %v \n", bookings)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v ", firstNames)

			noTickets := remainingTickets == 0
			if noTickets {
				fmt.Println("Our conference is booked out! Come next year ")
				break
			}

		} else {
			if !isValidEmail {
				fmt.Println("the email does not contain `@` ")
			}
			if !isValidName {
				fmt.Println("first name or last name is too short!")
			}
			if !isValidTicket {
				fmt.Println("number of input you entered is invalid")
			}
		}
	}

}

func greet(conferenceName string, totalTicket int, remainingTickets int) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("we have a total of %v tickets and %v are still available.\n", totalTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
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

	return firstName, lastName, userMail, userTickets

}

func validateUserInput(firstName string, lastName string, userMail string, userTickets int, remainingTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userMail, "@")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
