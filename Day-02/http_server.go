package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // Each request calls handler
	log.Fatal(http.ListenAndServe(":8090", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website! You've requested: %s", r.URL.Path)
}
