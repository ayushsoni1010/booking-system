package main

import "fmt"

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T.\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v  are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your ticket here to attend.\n")
	// ARRAYS
	// var bookings [50]string

	// SLICES
	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets int
	// ask user for their name

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

	remainingTickets = remainingTickets - uint(userTickets)
	// ARRAYS
	// bookings[0] = firstName + " " + lastName

	// SLICES
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("The whole array: %v\n", bookings)
	fmt.Printf("The first value: %v\n", bookings[0])
	fmt.Printf("The data type: %T\n", bookings)
	fmt.Printf("The aray length: %v\n", len(bookings))

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
