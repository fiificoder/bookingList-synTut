package main

import (
	"fmt"

	"github.com/fiificoder/synTut/validate"
)

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
