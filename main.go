package main

import (
	"fmt"
	"strings"
)

const noOfTickets int = 50

var conferenceName = "Go Conference"
var remaingTickets uint = 50 // uint - unsigned integer, only contains positive numbers. remaining tickets can't go beyond 0
// var attendees []string
// Another way of declaring variable, not applicable for const
var attendees = []string{}

func main() {

	var attendee string

	// greet the user with conference basic details
	greetUser()

	for {
		// take user input
		firstName, lastName, userEmail, userTickets := getUserInputs()

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, userEmail, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			attendee = firstName + " " + lastName
			// call book tickets
			bookTicket(userTickets, attendee, userEmail)

			attendees = append(attendees, attendee)

			// print first names of attendees
			printFirstNames()

			// check whether we ran out of tickets
			if remaingTickets == 0 {
				// end program
				fmt.Println("\nOur conference is booked out. Come back next year...")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("\nfirst name or last name you entered is too short...")
			}
			if !isValidEmail {
				fmt.Println("A valid email should contain one and only one @ symbol")
			}
			if !isValidTicketNumber {
				fmt.Printf("You have entered ticket count %d. We only have %d nos of tickets remaining...", userTickets, remaingTickets)
			}
		}
	}
}

// Greets user with conference name, total no of tickets and remaining tickets for the conference
func greetUser() {
	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}

// print first names of attendees
func printFirstNames() {
	// use of for-each loop
	firstNames := []string{}
	for _, attendee := range attendees {
		var names = strings.Fields(attendee) // spliting the full name from blank space
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("\nThe first names of attendees are : %v", firstNames)
}

// takes user input
func getUserInputs() (string, string, string, uint) {
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

func bookTicket(userTickets uint, attendee string, userEmail string) {
	// calculate remaining tickets after user booking
	remaingTickets -= userTickets

	fmt.Printf("Thank you %s for booking %d tickets. You will receive a notification email at %s", attendee, userTickets, userEmail)
	fmt.Printf("\nNumber of Tickets Available are %d", remaingTickets)
}
