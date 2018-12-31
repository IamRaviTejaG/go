package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book -> Structure containing the book attributes.
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2018"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2018"},
		Book{ID: 3, Title: "Golang routes", Author: "Mr. Router", Year: "2018"},
		Book{ID: 4, Title: "Golang in the building", Author: "Mr. Larry Page", Year: "2018"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
	log.Println("Gets all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idInt, _ := strconv.Atoi(params["id"])
	flag := 0
	for _, book := range books {
		if book.ID == idInt {
			flag = 1
			json.NewEncoder(w).Encode(&book)
			log.Println("Gets one book, ID: ", idInt)
		}
	}
	if flag == 0 {
		a := map[string]string{"IndexError": "Book ID not found!"}
		json.NewEncoder(w).Encode(&a)
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var newBookData Book
	json.NewDecoder(r.Body).Decode(&newBookData)
	books = append(books, newBookData)
	a := map[string]string{"Status": "Success"}
	json.NewEncoder(w).Encode(&a)
	log.Println("Added book: ", newBookData)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBookData Book
	flag := 0
	json.NewDecoder(r.Body).Decode(&updatedBookData)
	for index, item := range books {
		if item.ID == updatedBookData.ID {
			flag = 1
			books[index] = updatedBookData
			json.NewEncoder(w).Encode(&updatedBookData)
			log.Println("Updated book: ", updatedBookData)
		}
	}
	if flag == 0 {
		a := map[string]string{"IndexError": "Book ID not found!"}
		json.NewEncoder(w).Encode(&a)
		log.Println(a)
	}
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flag := 0
	id, _ := strconv.Atoi(params["id"])
	for i, item := range books {
		if item.ID == id {
			flag = 1
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode(&item)
			log.Println("Removed book: ", item)
		}
	}
	if flag == 0 {
		a := map[string]string{"IndexError": "Book ID not found!"}
		json.NewEncoder(w).Encode(&a)
		log.Println(a)
	}
}
