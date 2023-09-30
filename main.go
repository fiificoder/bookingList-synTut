package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 70
	var remainingTickets uint = 60

	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var bookings []string

	for {
		var firstName string
		var lastName string
		var userTicket uint
		var email string
		// ask user for their name
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your valid email:")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets:")
		fmt.Scan(&userTicket)

		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %s for booking %d tickets for the %s \n", firstName, userTicket, conferenceName)
		fmt.Printf("You will recieve an email confirmation to %v\n", email)
		fmt.Printf("%v tickets are now remaining for the conference\n", remainingTickets)

		var firstNames = []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Successful Bookings: %v \n", firstNames)
	}
}
