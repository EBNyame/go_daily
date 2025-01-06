package main

import (
	"fmt"
)

func passMe() {
	var username string
	var password string
	var lusername string
	var lpassword string

	// registration

	fmt.Println("Create a new Account...")
	fmt.Print("Enter your username: ")
	fmt.Scan(&username)

	fmt.Print("Enter your password: ")
	fmt.Scan(&password)

	// login
	fmt.Println("login your account")
	fmt.Println("Enter your username: ")
	fmt.Scan(&lusername)
	fmt.Print("Enter your password: ")
	fmt.Scan(&lpassword)

	// compare password and username
	if username == lusername && password == lpassword {
		fmt.Printf("Welcome! %s You've sucessfully login \n", lusername)
	} else {
		fmt.Println("Invalid credentials!")
	}

}
