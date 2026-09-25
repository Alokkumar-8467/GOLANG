package main

import (
	"fmt"
	"net/http"
)

func main() {
	url := "http://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)
