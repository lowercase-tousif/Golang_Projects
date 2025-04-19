package main

import "fmt"

func welcomeMessage() {
	fmt.Print("\n\n\t\t\t Even Odd Checker Game \t\t\t\t\n\n")
}

func evenOddChecker() {
	
	var num int
	
	fmt.Print("Enter any number: ")
	fmt.Scan(&num)

	if num % 2 == 0 {
		fmt.Print("\n\n\t\t\t\t Even Number ✨\t\t\t\t\n\n")
	} else{
		fmt.Print("\n\n\t\t\t\t Odd Number 🎉\t\t\t\t\n\n")
	}
}

func main() {
	welcomeMessage()
	evenOddChecker()
}
