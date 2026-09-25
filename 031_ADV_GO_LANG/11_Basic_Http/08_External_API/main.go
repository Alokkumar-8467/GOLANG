package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


type CatFactResponse struct {
	Fact   string `json:"fact"`
	Lenght string `json:"lenght"`
}


func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func fetchCatFact() (CatFactResponse, error) {
	url := "https://catfact.ninja/fact"

		res, err := http.Get(url)
	if err != nil {
		return CatFactResponse{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return CatFactResponse{}, fmt.Errorf("external api failed: %s", res.Status)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return CatFactResponse{}, err
	}
	


func externalHandler(w http.ResponseWriter, r *http.Request) {

}


func main() {

	http.HandleFunc("/external", externalHandler)

	err := http.ListenAndServe(":5000", nil)

	fmt.Println(err)

}
