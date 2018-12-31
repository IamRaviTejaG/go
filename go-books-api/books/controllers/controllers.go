package controllers

import (
	"books/dbcontrol"
	"books/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller struct
type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks returns all the books' details.
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		dbcontroller := dbcontrol.DBcontroller{}
		books = dbcontroller.GetBooks(db, book, books)
		json.NewEncoder(w).Encode(books)
	}
}

// GetBook returns a single book's details.
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		params := mux.Vars(r)
		dbcontroller := dbcontrol.DBcontroller{}
		book = dbcontroller.GetBook(db, book, params["id"])
		json.NewEncoder(w).Encode(book)
	}
}

// AddBook adds a book's details.
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int
		json.NewDecoder(r.Body).Decode(&book)
		dbcontroller := dbcontrol.DBcontroller{}
		bookID = dbcontroller.AddBook(db, book)
		json.NewEncoder(w).Encode(bookID)
	}
}

// UpdateBook updates a book's details.
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		json.NewDecoder(r.Body).Decode(&book)
		dbcontroller := dbcontrol.DBcontroller{}
		rowsUpdated := dbcontroller.UpdateBook(db, book)
		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

// RemoveBook removes a book's details.
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		dbcontroller := dbcontrol.DBcontroller{}
		rowsDeleted := dbcontroller.RemoveBook(db, params["id"])
		json.NewEncoder(w).Encode(rowsDeleted)
	}
}
