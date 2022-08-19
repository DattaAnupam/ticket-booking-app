package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const noOfTickets int = 50
	var remaingTickets int = 50

	fmt.Printf("Welcome to our %s booking application\n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d tickets are available\n", noOfTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend")
}
