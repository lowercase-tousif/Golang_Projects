package main

import (
	"student-management/handler"
	"student-management/message"
)

func main() {
	message.WelcomeMessage()
	handler.Menu()
}
