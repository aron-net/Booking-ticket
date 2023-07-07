package main

import (
	"fmt"
	"github.com/aron-helu/Booking-ticket/helper"
)

var conferenceName string = "Go Conference"

const conferenceTicket int = 50

var remainingTickets uint = 50

var bookings = []string{}

func main() {

	greetUser()

	// fmt.Printf("ConferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTicket, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	for {

		firstName, lastName, email, userTickets := getUserInput()

		// isValidCity := city == "Singapore" || city == "London"
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTickets(userTickets, firstName, lastName, email)
			// fmt.Printf("The whole array: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Array type: %T\n", bookings)
			// fmt.Printf("Array length: %v\n", len(bookings))

			// fmt.Printf("The whole Slice: %v\n", bookings)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			// var noTicketremaining bool = remainingTickets == 0
			// noTicketremaining := remainingTickets == 0
			// if noTicketremaining {
			// 	// do somthing
			// }

			if remainingTickets == 0 {
				//end
				fmt.Printf("Our comference is booked out. Come back next year. \n")
				break
			}
			// fmt.Println("Abrahm and Senait are planning their wedding")
		} else {
			// fmt.Printf("We only have %v tickets remaining, so you can't  book %v tickets.\n", remainingTickets, userTickets)
			if !isValidName {
				fmt.Println("First name or last name you entered to short")
			}
			if !isValidEmail {
				fmt.Println("The email address you entered does not contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("The number of tickets you entered is invalid;;")
			}
			fmt.Println("Your Input data is invalid try agin")
		}

	}

	fmt.Printf("These are all our bookings: %v\n", bookings)

	firstNames := printFirstName()

	fmt.Printf("The first names of of bookings: %v\n", firstNames)
	// Switch statement example
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")
	}
}

