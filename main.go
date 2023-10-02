package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 70

var remainingTickets uint = 60
var bookings []string

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTicket := UserInput()

		isValidName, isValidEmail, isValidTicket := InputValidation(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicket {
			bookTicket(userTicket, firstName, lastName, email)

			//list of firstNames of Users
			firstNames := listFirstNames()
			fmt.Printf("Successful Bookings(First names only): %v \n", firstNames)

			if remainingTickets == 0 {
				//end program
				fmt.Println(" ")
				fmt.Println("----------Sorry, You're late, TICKETS SOLD OUT!!.------------------ ")
				break
			}
		} else {
			if !isValidName {
				fmt.Printf("Invalid name. Try again\n")
			}
			if !isValidEmail {
				fmt.Printf("Invalid email. Try again\n")
			}
			if !isValidTicket {
				fmt.Printf("Invalid number of tickets. Available tickets : %v\n", remainingTickets)
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v  booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

}

func listFirstNames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func InputValidation(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicket := userTicket > 0 && userTicket <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}

func UserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTicket
}

func bookTicket(userTicket uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTicket
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %s for booking %d ticket(s) for the %s .\n", firstName, userTicket, conferenceName)
	fmt.Printf("Hello %s, you will recieve an email confirmation to %v .\n", firstName, email)
	fmt.Println(" ")
	fmt.Printf("Tickets remaining for conference: %v\n", remainingTickets)

}
