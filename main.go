package main

import (
	"fmt"
)

func main() {
	const noOfTickets int = 50

	var conferenceName string = "Go Conference"
	var remaingTickets uint = 50 // uint - unsigned integer, only contains positive numbers. remaining tickets can't go beyond 0
	var firstName, lastName, attendee, userEmail string
	var userTickets uint // uint - unsigned integer, only contains positive numbers. A user can't buy -1 nos tickets
	// var attendees []string
	// Another way of declaring variable, not applicable for const
	attendees := []string{}

	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		// take user input
		fmt.Println("\nEnter First name")
		fmt.Scanf("%s\n", &firstName)

		fmt.Println("Enter Last name")
		fmt.Scanf("%s\n", &lastName)

		fmt.Println("Enter email address")
		fmt.Scanf("%s\n", &userEmail)

		fmt.Println("Enter number of tickets you want to book")
		fmt.Scanf("%d\n", &userTickets)

		// calculate remaining tickets after user booking
		remaingTickets -= userTickets

		attendee = firstName + " " + lastName
		attendees = append(attendees, attendee)

		fmt.Printf("Thank you %s for booking %d tickets. You will receive a notification email at %s", attendee, userTickets, userEmail)

		fmt.Printf("\nNumber of Tickets Available are %d", remaingTickets)
		fmt.Printf("\nList of current attendees %v", attendees)
	}

}
