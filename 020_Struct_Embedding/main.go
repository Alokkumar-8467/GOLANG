package main

import (
	"fmt"
	"time"
)

// We can embed struct in another struct to use the fields of the embedded struct in the parent struct

// Like in Order struct we can
	status    string
	createdAt time.Time
	customer  // Embedded struct
}

func main() {

	newOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
		customer: customer{
