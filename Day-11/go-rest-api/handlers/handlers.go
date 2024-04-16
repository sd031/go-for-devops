package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"go-rest-api/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Handler struct {
	Collection *mongo.Collection
}

// GetBooks retrieves all books
func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []models.Book
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := h.Collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var book models.Book
		cursor.Decode(&book)
		books = append(books, book)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

// GetBook retrieves a single book by ID
func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := h.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// CreateBook creates a new book
func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := h.Collection.InsertOne(ctx, book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// UpdateBook updates an existing book
func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Book
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	json.NewDecoder(r.Body).Decode(&book)

	result, err := h.Collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{{"$set", book}},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// DeleteBook deletes a book
func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := h.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
