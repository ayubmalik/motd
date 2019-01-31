package main

import (
	"encoding/json"
	"net/http"
)

// Motd represents a Message of the Day
type Motd struct {
	Version string
	Message string
}

func motdHandler(w http.ResponseWriter, r *http.Request) {
	motd := Motd{"1.0.0", "hello world!"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(motd)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/motd", motdHandler)
	http.ListenAndServe(":8000", mux)
}
