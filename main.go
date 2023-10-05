package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const totalTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our %v application\n", conferenceName)
	fmt.Printf("we have total %v and %v are remaining\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int
	fmt.Println("Please write your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please write your Last name")
	fmt.Scan(&lastName)
	fmt.Println("Please write your email")
	fmt.Scan(&email)
	fmt.Println("Please tell us how many tickets you want")
	fmt.Scan(&userTickets)
	fmt.Printf("first name %v last name %v email %v and how many user want to buy tickets %v", firstName, lastName, email, userTickets)
}
