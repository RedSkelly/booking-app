package main

import "fmt"

func main() {

	conferenceName := "Go Conference" //the := can be used for declaring variables; cannot be used for constants, and cannot specify data type when using this shorthand
	const conferenceTickets = 50
	var remainingTickets uint = 50

	//what is each var's type?
	//fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to our %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets, and %v tickets are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend!\n")

	//store names of ticket purchasers in an array
	var bookings []string

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their name
	fmt.Printf("\nPlease enter your first name: \n")
	fmt.Scan(&firstName)

	fmt.Printf("\nPlease enter your last name: \n")
	fmt.Scan(&lastName)

	fmt.Printf("\nPlease enter your email address: \n")
	fmt.Scan(&email)

	fmt.Printf("\nPlease enter the number of tickets you would like to purchase: \n")
	fmt.Scan(&userTickets)
	
	remainingTickets = remainingTickets - userTickets
	//bookings[0] = firstName + " " + lastName
	bookings = append(bookings, firstName + " " + lastName)

	fmt.Printf("The whole slice: %v\n", bookings)
	fmt.Printf("The first element in array: %v\n", bookings[0])
	fmt.Printf("Slice type: %T\n", bookings)
	fmt.Printf("Slice length: %v\n", len(bookings))

	fmt.Printf("\nThank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

	fmt.Printf("There are now %v tickets remaining for the %v.\n", remainingTickets, conferenceName)


}

