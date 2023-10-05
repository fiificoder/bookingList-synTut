package main

import (
	"fmt"
	"time"

	"github.com/fiificoder/synTut/validate"
	_ "github.com/lib/pq"
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
	country         string
}

func main() {
	validate.ConnectDB()
	validate.GreetUsers(conferenceName, conferenceTickets, remainingTickets)

	for {
		firstName, lastName, email, userTicket, country := UserInput()

		isValidName, isValidEmail, isValidTicket := validate.InputValidation(firstName, lastName, email, userTicket, remainingTickets)

		if isValidName && isValidEmail && isValidTicket {

			bookTicket(userTicket, firstName, lastName, email, country)
			go sendTicket(firstName, lastName, userTicket, email)

			//list of firstNames of Users
			firstNames := listFirstNames()
			fmt.Printf("Successful Bookings(First names only): %v\n", firstNames)

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

func UserInput() (string, string, string, uint, string) {
	var firstName string
	var lastName string
	var userTicket uint
	var email string
	var country string
	// ask user for their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your valid email:")
	fmt.Scan(&email)

	fmt.Println("Enter Country of residence:")
	fmt.Scanln(&country)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTicket)

	_, err := validate.Db.Exec("INSERT INTO persons(firtname, lastname, email, country) VALUES($1, $2, $3, $4)", firstName, lastName, email, country)
	if err != nil {
		fmt.Println("Error storing user data, Please try again")
	}

	return firstName, lastName, email, userTicket, country
}

func bookTicket(userTicket uint, firstName string, lastName string, email string, country string) {
	remainingTickets = remainingTickets - userTicket

	userData := UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTicket,
		country:         country,
	}

	_, err := validate.Db.Exec("INSERT INTO persons (firstname, lastname, email, usertickets, country) VALUES($1, $2, $3, $4, $5)", userData.firstName, userData.lastName, userData.email, userData.numberOfTickets, userData.country)
	if err != nil {
		fmt.Println("Error storing user data, Please try again")
	}

	bookings = append(bookings, userData)

	//fmt.Printf("List of bookings is: %v\n", userData)

	fmt.Printf("Thank you %s for booking %d ticket(s) for the %s .\n", firstName, userTicket, conferenceName)
	fmt.Printf("Hello %s, you will recieve an email confirmation to %v .\n", firstName, email)
	fmt.Println(" ")
	fmt.Printf("Tickets remaining for conference: %v\n", remainingTickets)

}

func sendTicket(firstName string, lastName string, userTicket uint, email string) {
	time.Sleep(40 * time.Second)
	ticket := fmt.Sprintf("%v tickets for %v %v \n", userTicket, firstName, lastName)
	fmt.Println("##################################")
	fmt.Printf("Sending tickets:\n %v\nto email address: %v\n", ticket, email)
	fmt.Println("##################################")
}
