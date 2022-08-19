package main

import (
	"fmt"
)

func main() {
	const noOfTickets int = 50

	var conferenceName string = "Go Conference"
	var remaingTickets int = 50
	var firstName, lastName, fullName, userEmail string
	var userTickets int

	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")

	// take user input
	fmt.Println("Enter First name")
	fmt.Scanf("%s\n", &firstName)

	fmt.Println("Enter Last name")
	fmt.Scanf("%s\n", &lastName)

	fmt.Println("Enter email address")
	fmt.Scanf("%s\n", &userEmail)

	fmt.Println("Enter number of tickets you want to book")
	fmt.Scanf("%d\n", &userTickets)

	fullName = firstName + " " + lastName
	fmt.Printf("Thank you %s for booking %d tickets. You will receive a notification email at %s", fullName, userTickets, userEmail)
}
