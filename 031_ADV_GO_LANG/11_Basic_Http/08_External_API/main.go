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


func main() {

}
