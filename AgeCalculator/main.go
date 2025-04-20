package main

import (
	"fmt"
	"time"
)

// Welcome Message
func welcomeMessage() {
	fmt.Print("\n\n\t\t\t\t Age Calculator 📆\t\t\t\t\n\n")
}

func userInput() {

	// Getting the user input separately like year , month and date
	fmt.Print("Enter your birth year: ")
	var userYear int
	fmt.Scan(&userYear)

	fmt.Print("Enter your birth Month: ")
	var userMonth int
	fmt.Scan(&userMonth)

	fmt.Print("Enter your birth date: ")
	var userDay int
	fmt.Scan(&userDay)

	// Calling the calcYear Function
	calcYear(userYear, userMonth, userDay)
}

func calcYear(userYear int, userMonth int, userDay int) {

	// getting the current year using time.Now().Year()
	var currentYear int = time.Now().Year()
	var currentMonth int = int(time.Now().Month())
	var currentDay int = time.Now().Day()

	// Calculating the result for the year
	var result int = currentYear - userYear

	// Checking if the birthday appread this year or not
	if (currentMonth < userMonth) || (currentMonth == userMonth && (currentDay < userDay)) {
		result--
	}

	fmt.Printf("\n\n\t\t\t\t You are %d years old.\t\t\t\t\n\n", result)
}

func main() {
	welcomeMessage()
	userInput()
}
