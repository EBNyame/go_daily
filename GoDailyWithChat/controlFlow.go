package main

// Exercise: Control Structures Practice
// Write a program that checks if a number is even or odd using an if statement.
// Create a loop that prints all even numbers from 1 to 20.
// Given an array of numbers, use a loop to find and print the sum of all elements.

func check(a int) string {
	if a <= 0 {
		return "Invalid"
	} else if a%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
