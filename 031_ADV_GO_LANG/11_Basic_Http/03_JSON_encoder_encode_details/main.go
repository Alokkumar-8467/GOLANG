package main

import (
	"fmt"
	"net/http"
)

func successHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

func main() {

	http.HandleFunc("/ok", successHandler)

	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)

}
