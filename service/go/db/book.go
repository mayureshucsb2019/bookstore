package db

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Book struct represents the structure of a book record in the database.
type Book struct {
	ISBN            string
	Name            string
	Tags            []string
	AuthorName      string
	DateOfPublish   string
	PublishingHouse string
	NumberOfPages   int
	Cost            float64
}

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// BookRepository provides access to the book storage.
type BookRepository struct {
	DB *sql.DB
}

// NewBookRepository creates a new BookRepository with a database connection.
func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

// InitDB initializes and returns a MySQL database connection.
func InitDB(config DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// CreateBook inserts a new book into the database.
func (r *BookRepository) CreateBook(book *Book) error {
	query := `INSERT INTO Books (isbn, name, tags, author_name, date_of_publish, publishing_house, number_of_pages, cost) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.DB.Exec(query, book.ISBN, book.Name, fmt.Sprintf("%v", book.Tags), book.AuthorName, book.DateOfPublish, book.PublishingHouse, book.NumberOfPages, book.Cost)
	return err
}

// GetBookByISBN retrieves a book from the database by its ISBN.
func (r *BookRepository) GetBookByISBN(isbn string) (*Book, error) {
	query := `SELECT * FROM Books WHERE isbn = ?`
	row := r.DB.QueryRow(query, isbn)

	var book Book
	var tags string

	err := row.Scan(&book.ISBN, &book.Name, &tags, &book.AuthorName, &book.DateOfPublish, &book.PublishingHouse, &book.NumberOfPages, &book.Cost)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // No book found
		}
		return nil, err
	}

	// Convert tags from string to slice
	book.Tags = parseTags(tags)

	return &book, nil
}

// UpdateBook updates an existing book record in the database.
func (r *BookRepository) UpdateBook(book *Book) error {
	query := `UPDATE Books SET name=?, tags=?, author_name=?, date_of_publish=?, publishing_house=?, number_of_pages=?, cost=? WHERE isbn=?`
	_, err := r.DB.Exec(query, book.Name, fmt.Sprintf("%v", book.Tags), book.AuthorName, book.DateOfPublish, book.PublishingHouse, book.NumberOfPages, book.Cost, book.ISBN)
	return err
}

// DeleteBook removes a book from the database by its ISBN.
func (r *BookRepository) DeleteBook(isbn string) error {
	query := `DELETE FROM Books WHERE isbn = ?`
	_, err := r.DB.Exec(query, isbn)
	return err
}

// GetAllBooks retrieves all books from the database.
func (r *BookRepository) GetAllBooks() ([]Book, error) {
	rows, err := r.DB.Query("SELECT * FROM Books")
	if err != nil {
		return nil, fmt.Errorf("failed to query books: %w", err)
	}
	defer rows.Close()

	var books []Book
	var tags string
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ISBN, &book.Name, &tags, &book.AuthorName, &book.DateOfPublish, &book.PublishingHouse, &book.NumberOfPages, &book.Cost); err != nil {
			return nil, fmt.Errorf("failed to scan book: %w", err)
		}
		// Convert tags from string to slice
		book.Tags = parseTags(tags)
		books = append(books, book)
	}

	return books, nil
}

// Helper function to parse tags from a string to a slice.
func parseTags(tags string) []string {
	// Implement the logic to convert the tags string back to a slice
	return []string{"Check"} // Placeholder: implement your tag parsing logic
}
