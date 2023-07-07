package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50

	bookings := []string{}

	fmt.Printf("ConferenceTicket is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		// fmt.Printf("The whole array: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Array type: %T\n", bookings)
		// fmt.Printf("Array length: %v\n", len(bookings))

		// fmt.Printf("The whole Slice: %v\n", bookings)
		// fmt.Printf("The first value: %v\n", bookings[0])
		// fmt.Printf("Slice type: %T\n", bookings)
		// fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining. \n", remainingTickets)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames =	append(firstNames, names[0])
		}
		fmt.Printf("These are all our bookings: %v\n", bookings)
		fmt.Printf("The first names of of bookings: %v\n", firstNames)

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
	}

}
