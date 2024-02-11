package main

import (
	"log"
	"net/http"

	"go-rest-api/db"
	"go-rest-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Connect to database
	collection := db.ConnectDB()

	// Create an instance of Handler with the collection
	handler := &handlers.Handler{Collection: collection}

	// Setup routes
	router.HandleFunc("/books", handler.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
	router.HandleFunc("/books", handler.CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
