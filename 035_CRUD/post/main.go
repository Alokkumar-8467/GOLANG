package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
func performGetRequest() {
	url := "http://jsonplaceholder.typicode.com/todos/1"

	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response: ", resp.Status)
		return
	}
	// **************
	// Advanced method to get data from URL
	var todo Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON: ", err)
		return
	}
	fmt.Println("Decoded Todo: ", todo)
}
