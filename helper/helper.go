package helper

import (
	"fmt"
	"strings"
)

type UserData struct {
	FirstName   string
	LastName    string
	UserEmail   string
	UserTickets uint
}

type ValidationTypes struct {
	IsValidName         bool
	IsValidEmail        bool
	IsValidTicketNumber bool
}

// Greets user with conference name, total no of tickets and remaining tickets for the conference
func GreetUser(conferenceName string, noOfTickets int, remaingTickets uint) {
	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}

// validates inputs from the user
func ValidateUserInput(userData UserData, remaingTickets uint) (bool, bool, bool) {
	// firstName & lastName should contain atleast  characters
	isValidName := len(userData.FirstName) >= 2 && len(userData.LastName) >= 2
	// email should contain @ character
	isValidEmail := strings.Contains(userData.UserEmail, "@")
	// user should book atleast 1 ticket and it should be less remainTickets
	isValidTicketNumber := userData.UserTickets > 0 && userData.UserTickets <= remaingTickets

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

func ShowValidationErrMsg(validationType ValidationTypes, userTickets uint, remaingTickets uint) {
	if !validationType.IsValidName {
		fmt.Println("\nfirst name or last name you entered is too short...")
	}
	if !validationType.IsValidEmail {
		fmt.Println("A valid email should contain one and only one @ symbol")
	}
	if !validationType.IsValidTicketNumber {
		fmt.Printf("You have entered ticket count %d. We only have %d nos of tickets remaining...", userTickets, remaingTickets)
	}
}
