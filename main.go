package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) < 50 {
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

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketCount := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketCount {
			remainingTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("There are %v remaining tickets for %v.\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets <= 0 {
				// end program
				fmt.Println("Our conference is fully booked. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Your first name or last name is invalid. Try again.")
			}
			if !isValidEmail {
				fmt.Println("Your email is invalid. Try again.")
			}
			if !isValidTicketCount {
				fmt.Println("Your purchasing too many or too little tickets. Try again.")
			}
		}
	}
}

func greetUser(confName string, confTickets int, remTickets int) {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}
