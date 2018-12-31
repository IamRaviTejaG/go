package dbcontrol

import (
	"books/models"
	"database/sql"
	"log"
)

// DBcontroller struct
type DBcontroller struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetBooks returns slice of books
func (d DBcontroller) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("SELECT * FROM books")
	logFatal(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)
		books = append(books, book)
	}
	return books
}

// GetBook returns info of one book
func (d DBcontroller) GetBook(db *sql.DB, book models.Book, id string) models.Book {
	rows := db.QueryRow("SELECT * FROM books WHERE id=$1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)
	return book
}

// AddBook adds book to DB
func (d DBcontroller) AddBook(db *sql.DB, book models.Book) int {
	err := db.QueryRow("INSERT INTO books(title, author, year) VALUES($1, $2, $3) RETURNING id",
		book.Title, book.Author, book.Year).Scan(&book.ID)
	logFatal(err)
	return book.ID
}

// UpdateBook updates the details of pre-existing record
func (d DBcontroller) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	logFatal(err)
	rowsUpdated, err := result.RowsAffected()
	logFatal(err)
	return rowsUpdated
}

// RemoveBook removes book entry from the DB
func (d DBcontroller) RemoveBook(db *sql.DB, id string) int64 {
	result, err := db.Exec("DELETE FROM books WHERE id=$1", id)
	logFatal(err)
	rowsDeleted, err := result.RowsAffected()
	logFatal(err)
	return rowsDeleted
}
