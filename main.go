package main

import (
	"fmt"
	"main/helper"
	"strconv"
)

const noOfTickets int = 50

var conferenceName = "Go Conference"
var remaingTickets uint = 50 // uint - unsigned integer, only contains positive numbers. remaining tickets can't go beyond 0
// var attendees []string
// Another way of declaring variable, not applicable for const
var attendees = make([]map[string]string, 0)

func main() {

	var attendee string

	// greet the user with conference basic details
	helper.GreetUser(conferenceName, noOfTickets, remaingTickets)

	for {
		// take user input
		firstName, lastName, userEmail, userTickets := helper.GetUserInputs()

		// create a map for a user
		var userData = make(map[string]string)
		userData["firstName"] = firstName
		userData["lastName"] = lastName
		userData["userEmail"] = userEmail
		userData["userTickets"] = strconv.FormatUint(uint64(userTickets), 10) // convert uint to string

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userData, remaingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			attendee = userData["firstName"] + " " + userData["lastName"]
			// call book tickets
			bookTicket(userData, attendee)

			attendees = append(attendees, userData)

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

// print first names of attendees
func printFirstNames() {
	// use of for-each loop
	firstNames := []string{}
	for _, attendee := range attendees {
		firstNames = append(firstNames, attendee["firstName"])
	}
	fmt.Print("\nThe first names of attendees are :")
	for _, fName := range firstNames {
		fmt.Printf("%s ", fName)
	}
}

func bookTicket(userData map[string]string, attendee string) {
	// calculate remaining tickets after user booking
	userTickets, _ := strconv.Atoi(userData["userTickets"])
	remaingTickets -= uint(userTickets)

	fmt.Printf("Thank you %s for booking %d tickets. You will receive a notification email at %s", attendee, userTickets, userData["userEmail"])
	fmt.Printf("\nNumber of Tickets Available are %d", remaingTickets)
}
