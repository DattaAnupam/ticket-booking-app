package helper

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type UserData struct {
	FirstName   string
	LastName    string
	UserEmail   string
	UserTickets uint8
}

type ValidationTypes struct {
	IsValidName         bool
	IsValidEmail        bool
	IsValidTicketNumber bool
}

// Greets user with conference name, total no of tickets and remaining tickets for the conference
func GreetUser(conferenceName string, noOfTickets uint8, remaingTickets uint8) {
	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}

// validates inputs from the user
func ValidateUserInput(userData UserData, remaingTickets uint8) (bool, bool, bool) {
	// firstName & lastName should contain atleast  characters
	isValidName := len(userData.FirstName) >= 2 && len(userData.LastName) >= 2
	// email should contain @ character
	isValidEmail := strings.Contains(userData.UserEmail, "@")
	// user should book atleast 1 ticket and it should be less remainTickets
	isValidTicketNumber := userData.UserTickets > 0 && userData.UserTickets <= remaingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// takes user input
func GetUserInputs() (string, string, string, uint8) {
	// var userTickets uint // uint - unsigned integer, only contains positive numbers. A user can't buy -1 nos tickets

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nEnter First name: ")
	firstName, _ := reader.ReadString('\n')  // take input until pressed enter, i.e. new line
	firstName = strings.TrimSpace(firstName) // Trim white spaces

	fmt.Print("Enter Last name: ")
	lastName, _ := reader.ReadString('\n')
	lastName = strings.TrimSpace(lastName)

	fmt.Print("Enter email address: ")
	userEmail, _ := reader.ReadString('\n')
	userEmail = strings.TrimSpace(userEmail)

	fmt.Print("Enter number of tickets you want to book: ")
	tickets, _ := reader.ReadString('\n')
	tickets = strings.TrimSpace(tickets)
	userTickets, _ := strconv.ParseUint(tickets, 10, 64)

	return firstName, lastName, userEmail, uint8(userTickets)
}

func ShowValidationErrMsg(validationType ValidationTypes, userTickets uint8, remaingTickets uint8) {
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
