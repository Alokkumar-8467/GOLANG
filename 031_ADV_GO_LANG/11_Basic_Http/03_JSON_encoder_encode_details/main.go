package main

import (
	"fmt"
	"net/http"
)

func successHandler()

func main() {

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)

}
