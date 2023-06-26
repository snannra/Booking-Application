package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name
	fmt.Print("Type in your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Type in your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Type in your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of tickets you will buy: ")
	fmt.Scan(&userTickets)

	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
	fmt.Printf("There are %v remaining tickets for %v.\n", remainingTickets, conferenceName)

}
