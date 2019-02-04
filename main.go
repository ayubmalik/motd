package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Motd represents a Message of the Day
type Motd struct {
	Version string
	Message string
}

func motdHandler(w http.ResponseWriter, r *http.Request) {
	motd := Motd{"1.0.1", "hello world!"}
	w.Header().Set("Content-Type", "application/json")
	log.Printf("motd: %v", motd)
	json.NewEncoder(w).Encode(motd)
}

func main() {
	log.SetOutput(os.Stdout)
	mux := http.NewServeMux()
	mux.HandleFunc("/motd", motdHandler)
	http.ListenAndServe(":8000", mux)
}
