package main

import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome try to /hello?name=alok"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {


func main() {

	err := http.ListenAndServe()

}
