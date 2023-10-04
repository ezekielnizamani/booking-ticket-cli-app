package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const totalTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to our", conferenceName, "application")
	fmt.Println("we have total ", totalTickets, "and ", remainingTickets, "are remaining")
	fmt.Println("Get your tickets to attend")
}
