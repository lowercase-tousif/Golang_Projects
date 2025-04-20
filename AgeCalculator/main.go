package main

import (
	"fmt"
	"time"
)

// Welcome Message
func welcomeMessage() {
	fmt.Print("\n\n\t\t\t\t Age Calculator 📆\t\t\t\t\n\n")
}

// Calculate The BirthYear
func calcYear() {
	fmt.Print("Enter your birth year: ")

	var userYear int
	fmt.Scan(&userYear)

	// getting the current year using time.Now().Year()
	var currentYear int = time.Now().Year()

	var result int = currentYear - userYear

	fmt.Printf("\n\n\t\t\t\t You are %d years old.\t\t\t\t\n\n", result)

}
func main() {
	welcomeMessage()
	calcYear()
}
