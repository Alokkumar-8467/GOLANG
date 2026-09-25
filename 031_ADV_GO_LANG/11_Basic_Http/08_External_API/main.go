package main

import (
	"fmt"
	"net/http"
)



type CatFactResponse struct {
	Fact   string `json:"fact"`
	Lenght string `json:"lenght"`
}


func writeJSON(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}


func externalHandler(w http.ResponseWriter, r *http.Request) {

}


func main() {

	http.HandleFunc("/external", externalHandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)

}
