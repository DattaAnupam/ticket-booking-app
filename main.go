package main

import (
	"fmt"
	"main/helper"
)

const noOfTickets int = 50

var conferenceName = "Go Conference"
var remaingTickets uint = 50 // uint - unsigned integer, only contains positive numbers. remaining tickets can't go beyond 0
// var attendees []string
// Another way of declaring variable, not applicable for const
var attendees = []helper.UserData{}

func main() {

	var attendee string

	// greet the user with conference basic details
	helper.GreetUser(conferenceName, noOfTickets, remaingTickets)

	for {
		// take user input
		firstName, lastName, userEmail, userTickets := helper.GetUserInputs()

		// create a struct for a user
		var userData = helper.UserData{
			FirstName:   firstName,
			LastName:    lastName,
			UserEmail:   userEmail,
			UserTickets: userTickets,
		}

		// user input validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(userData, remaingTickets)

		// create a struct for validation types
		var validationTypes = helper.ValidationTypes{
			IsValidName:         isValidName,
			IsValidEmail:        isValidEmail,
			IsValidTicketNumber: isValidTicketNumber,
		}

		if isValidName && isValidEmail && isValidTicketNumber {
			attendee = userData.FirstName + " " + userData.LastName
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
			helper.ShowValidationErrMsg(validationTypes, userData.UserTickets, remaingTickets)
		}
	}
}

// print first names of attendees
func printFirstNames() {
	// use of for-each loop
	firstNames := []string{}
	for _, attendee := range attendees {
		firstNames = append(firstNames, attendee.FirstName)
	}
	fmt.Print("\nThe first names of attendees are :")
	for _, fName := range firstNames {
		fmt.Printf("%s ", fName)
	}
}

func bookTicket(userData helper.UserData, attendee string) {
	// calculate remaining tickets after user booking
	remaingTickets -= userData.UserTickets

	fmt.Printf("Thank you %s for booking %d tickets. You will receive a notification email at %s", attendee, userData.UserTickets, userData.UserEmail)
	fmt.Printf("\nNumber of Tickets Available are %d", remaingTickets)
}
