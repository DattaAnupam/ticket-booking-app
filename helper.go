package main

import (
	"strings"
)

// validates inputs from the user
func validateUserInput(firstName string, lastName string, userEmail string, userTickets uint) (bool, bool, bool) {
	// firstName & lastName should contain atleast  characters
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	// email should contain @ character
	isValidEmail := strings.Contains(userEmail, "@")
	// user should book atleast 1 ticket and it should be less remainTickets
	isValidTicketNumber := userTickets > 0 && userTickets <= remaingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
