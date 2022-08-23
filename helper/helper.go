package helper

import (
	"fmt"
	"strconv"
	"strings"
)

// Greets user with conference name, total no of tickets and remaining tickets for the conference
func GreetUser(conferenceName string, noOfTickets int, remaingTickets uint) {
	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}

// validates inputs from the user
func ValidateUserInput(userData map[string]string, remaingTickets uint) (bool, bool, bool) {
	// firstName & lastName should contain atleast  characters
	isValidName := len(userData["firstName"]) >= 2 && len(userData["lastName"]) >= 2
	// email should contain @ character
	isValidEmail := strings.Contains(userData["userEmail"], "@")
	// user should book atleast 1 ticket and it should be less remainTickets
	userTickets, _ := strconv.Atoi(userData["userTickets"])
	isValidTicketNumber := userTickets > 0 && uint(userTickets) <= remaingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// takes user input
func GetUserInputs() (string, string, string, uint) {
	var firstName, lastName, userEmail string
	var userTickets uint // uint - unsigned integer, only contains positive numbers. A user can't buy -1 nos tickets

	fmt.Println("\nEnter First name")
	fmt.Scanf("%s\n", &firstName)

	fmt.Println("Enter Last name")
	fmt.Scanf("%s\n", &lastName)

	fmt.Println("Enter email address")
	fmt.Scanf("%s\n", &userEmail)

	fmt.Println("Enter number of tickets you want to book")
	fmt.Scanf("%d\n", &userTickets)

	return firstName, lastName, userEmail, userTickets
}
