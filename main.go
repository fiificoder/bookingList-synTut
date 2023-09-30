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

		if userTicket > remainingTickets {
			fmt.Printf("Sorry, Only %v tickets are availabe now, so you can't book %v tickets .!!!!!\n", remainingTickets, userTicket)
			break
		}

		remainingTickets = remainingTickets - userTicket
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %s for booking %d ticket(s) for the %s .\n", firstName, userTicket, conferenceName)
		fmt.Printf("Hello %s, you will recieve an email confirmation to %v .\n", firstName, email)
		fmt.Println(" ")
		fmt.Printf("Tickets remaining for conference: %v\n", remainingTickets)

		var firstNames = []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("Successful Bookings: %v \n", firstNames)

		if remainingTickets == 0 {
			//end program
			fmt.Println(" ")
			fmt.Println("----------Sorry, You're late, TICKETS SOLD OUT!!.------------------ ")
			break
		}

	}
}
