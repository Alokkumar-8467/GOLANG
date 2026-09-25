package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func writeJSON(w http.ResponseWriter, status int, data any) {
		/*
	   The response we send back (w) needs a "Content-Type" header, so the
	   client knows the body is JSON and can parse it correctly.

	   To set that header, we first need the header collection of the
	   response, which we get using w.Header(). That returned object has
	   a .Set(key, value) method on it, which we use to actually set the
	   Content-Type header.

	   So the two steps combined into one line:
	       w.Header().Set("Content-Type", "application/json")
	*/
		w.Header().Set("Content-Type", "application/json")

	/* `w.WriteHeader(status)`
	This does two things at once, and both matter:
	1. It sets the HTTP status code
	2. It sends/flushes everything staged so far

	w.Header().Set("Content-Type", "application/json")   // staged, not sent yet
	w.WriteHeader(status)                                  // ← NOW it's actually sent
	*/
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

type TestRequest struct {
	Name string `json:"name"`
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    "false",
			"error": "Only post is allowed",
		})
		return
	}

	defer r.Body.Close()

	var req TestRequest

	dec := json.NewDecoder(r.Body)

	if err := dec.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    "false",
			"error": "Invalid json format",
		})
		return
	}

	// Now validation check
	req.Name = strings.TrimSpace(req.Name)

	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"ok":    "false",
			"error": "Name must not be empty",
		})
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"Ok":        "true",
		"data":      req,
		"timeStamp": time.Now().UTC(),
	})

}

func main() {

	http.HandleFunc("/test", testHandler)

	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)

}
