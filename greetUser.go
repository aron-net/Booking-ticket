package main

import "fmt"

func greetUser() {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Printf("We have the total of %v tickets and %v are still available\n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

