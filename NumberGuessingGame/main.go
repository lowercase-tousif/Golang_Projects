package main

import "fmt"

// Welcome Message
func welcomeMessage() {
	fmt.Print("\n\n\t\t\t Welcome to the Simple Number Guessing Game\t\t\t\n\n")
}


// Main Game Function
func mainGameLoop() {

	var attempt int = 0
	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scan(&guess)

		attempt++

		var secretNumber = ...

		if guess < secretNumber {
			fmt.Println("Too low try again!")
		} else if guess > secretNumber {
			fmt.Println("Too high try again!")
		} else {
			fmt.Println("You have successfully guessed the number✨")
			break;
		}

		if attempt == 3{
			fmt.Println("Game Ended")
			break;
		}
	}
}

func main() {
	welcomeMessage()
}
