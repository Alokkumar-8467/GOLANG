package main

import (
	"fmt"
	"time"
)

// Structs are basically custom Data Structure

// Order Struct

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
}

func main() {

	myOrder := order{
		id:     "1",
		amount: 50.00,
		status: "received",
	}

}
