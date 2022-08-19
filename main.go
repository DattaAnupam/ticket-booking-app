package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const noOfTickets int = 50
	var remaingTickets int = 50

	fmt.Println("Welcome to our", conferenceName, "booking application")
	fmt.Println("We have total of", noOfTickets, "tickets and", remaingTickets, "tickets are available")
	fmt.Println("Get your tickets here to attend")
}
