package main

import (
	r "firstpayment/router"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

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
	fmt.Print("Fist App")
	router := r.RouterConfig()

	log.Fatal(http.ListenAndServe(":8080", router))
}
