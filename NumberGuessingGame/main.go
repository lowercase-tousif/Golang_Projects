package main

import (
	"fmt"
	"math/rand"
)

// Welcome Message
func welcomeMessage() {
	fmt.Print("\n\n\t\t\t Welcome to the Simple Number Guessing Game\t\t\t\n\n")
	fmt.Print("\t\t\t You have only 3 guesses left\t\t\t\n\n")
	fmt.Print("\t\t\t Number is between 1-10 \t\t\t\n\n")

}

// Secret Number Generator
func generateSecretNumber() int{
	var secretNumber int 

	secretNumber = rand.Intn(11) + 1
	return secretNumber
}

// Main Game Function
func mainGameLoop() {

	var attempt int = 0
	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scan(&guess)

		attempt++

		var secretNumber int = generateSecretNumber()

		if guess < secretNumber {
			fmt.Println("Too low try again!")
		} else if guess > secretNumber {
			fmt.Println("Too high try again!")
		} else {
			fmt.Print("\n\n\t\t\t\t You have successfully guessed the number✨\t\t\t\t\n\n")
			break;
		}

		if attempt == 3{
			fmt.Print("\n\n\t\t\t\t ❌❌ Game Over ❌❌\t\t\t\t\n\n")
			break;
		}
	}
}

func main() {
	welcomeMessage()
	mainGameLoop()
}
