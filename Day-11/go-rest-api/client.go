package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-rest-api/models"
)

type Book struct {
	ID     string `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
	ISBN   string `json:"ISBN"`
}

func main() {

	// Base URL of your REST API
	baseURL := "http://localhost:8000"

	// Sample Book
	sampleBook := models.Book{
		Title:  "Sample Book",
		Author: "John Doe",
		ISBN:   "123-456-789",
	}

	// Create a new book
	createBook(baseURL, sampleBook)

	// Get all books
	books, err := getBooks(baseURL)

	if err != nil {
		// handle error
	}

	for _, book := range books {
		fmt.Printf("Book: %+v\n", book)
	}

	fmt.Println("Books=", books)
	fmt.Println("First Book id=", books[0].ID)

	// Update a book (replace {id} with actual book ID)
	sampleBook.Title = "Updated Book"
	updateBook(baseURL, books[0].ID, sampleBook)

	// Get a single book (replace {id} with actual book ID)
	getBook(baseURL, books[0].ID)

	// Delete a book (replace {id} with actual book ID)
	deleteBook(baseURL, books[0].ID)
}

func createBook(baseURL string, book models.Book) {
	jsonBook, _ := json.Marshal(book)
	resp, err := http.Post(baseURL+"/books", "application/json", bytes.NewBuffer(jsonBook))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Create Book Response:", string(body))
}

func getBooks(baseURL string) ([]Book, error) {
	resp, err := http.Get(baseURL + "/books")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("Get Books Response:", string(body))

	var books []Book
	jsonErr := json.NewDecoder(resp.Body).Decode(&books)
	if err != nil {
		fmt.Println("Error decoding JSON:", jsonErr)
		return nil, jsonErr
	}

	return books, nil
}

func getBook(baseURL, id string) {
	resp, err := http.Get(baseURL + "/books/" + id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Get Book Response:", string(body))
}

func updateBook(baseURL, id string, book models.Book) {
	jsonBook, _ := json.Marshal(book)
	req, _ := http.NewRequest(http.MethodPut, baseURL+"/books/"+id, bytes.NewBuffer(jsonBook))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Update Book Response:", string(body))
}

func deleteBook(baseURL, id string) {
	req, _ := http.NewRequest(http.MethodDelete, baseURL+"/books/"+id, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Delete Book Response:", string(body))
}
