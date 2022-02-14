package main

import "time"

// card
// user
type User struct {
	firstName   string
	middleName  string
	lastName    string
	cardDetails card
}

type card struct {
	cardNum      string
	cardName     string
	cvv          int
	expiryDate   time.Time
	cardProvider string
	issuerBank   string
}

func main() {

}
