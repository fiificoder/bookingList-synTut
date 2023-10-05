package main

import (
	"fmt"

	"github.com/fiificoder/synTut/validate"
)

var conferenceName = "Go Conference"

const conferenceTickets = 70

var remainingTickets uint = 60
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	validate.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTicket := UserInput()

		isValidName, isValidEmail, isValidTicket := validate.InputValidation(firstName, lastName, email, userTicket, remainingTickets)

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

func listFirstNames() []string {
	var firstNames = []string{}
	for _, booking := range bookings {
		//var names = strings.Fields(booking)
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
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

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is: %v\n", userData)

	fmt.Printf("Thank you %s for booking %d ticket(s) for the %s .\n", firstName, userTicket, conferenceName)
	fmt.Printf("Hello %s, you will recieve an email confirmation to %v .\n", firstName, email)
	fmt.Println(" ")
	fmt.Printf("Tickets remaining for conference: %v\n", remainingTickets)

}
