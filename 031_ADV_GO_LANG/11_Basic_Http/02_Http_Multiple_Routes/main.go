package main

import "net/http"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Welcome try to /hello?name=alok"))
}

func main() {

	err := http.ListenAndServe()

}
