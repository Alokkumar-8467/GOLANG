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
