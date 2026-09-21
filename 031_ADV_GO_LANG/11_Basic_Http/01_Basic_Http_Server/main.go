package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Response){
	r.Method  != http.MethodGet {
		http.Error(w, "only Get is allowed", http.StatusMethodNotAllowed)
		return
	}

	_, _ = w.write([]byte("Hello from GO net/http server"))

}

func main() {

	// http
	// How to register a router.
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Try going to 8000 port")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)

}
