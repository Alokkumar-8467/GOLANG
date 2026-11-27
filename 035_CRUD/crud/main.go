package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/todos"
	url1 := "http://jsonplaceholder.typicode.com/todos/1"

	res, err := http.Get(url)
	resp, err := http.Get(url1)

	if err != nil {
		fmt.Println("Error getting : ", err)
		return
	}
