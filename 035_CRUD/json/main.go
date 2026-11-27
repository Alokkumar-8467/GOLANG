package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	person := Person{
		Name:    "John Doe",
		Age:     30,
		IsAdult: true,
	}
	fmt.Println("Direct Structure Data", person)

	// Convert person JSON Encoding (Marshal)
	// Means converting struct into JSON data

	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return

	}
