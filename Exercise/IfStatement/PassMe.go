package main

import "fmt"

func PassMe(){
	func create(){
		var username, password string
		fmt.Println("Enter your username:")
		fmt.Scan(&username)
		fmt.Println("Enter your password:")
		fmt.Scan(&password)
	}
}