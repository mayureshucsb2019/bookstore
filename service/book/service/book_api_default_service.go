package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/mayureshucsb2019/bookstore/service/book/db"
	"github.com/mayureshucsb2019/bookstore/service/book/models"
	"github.com/mayureshucsb2019/bookstore/service/common"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPI API.
// This service interacts with the repository layer for data access.
type DefaultAPIService struct {
	Repo *db.BookRepository // Add a field to hold the repository
}

// NewDefaultAPIService creates a default API service with the given repository.
func NewDefaultAPIService(repo *db.BookRepository) *DefaultAPIService {
	return &DefaultAPIService{
		Repo: repo,
	}
}

// BooksGet - Get a paginated list of books
func (s *DefaultAPIService) BooksGet(ctx context.Context, pageNumber int32, pageSize int32) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	books, err := s.Repo.GetAllBooks() // Use the repository to get the books
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}
	var booksJSON []map[string]interface{}
	for _, book := range books {
		booksJSON = append(booksJSON, convertBookToAPIFormat(book))
	}

	return common.Response(http.StatusOK, booksJSON), nil
}

// BooksIsbnDelete - Delete a book by ISBN
func (s *DefaultAPIService) BooksIsbnDelete(ctx context.Context, isbn string) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	err := s.Repo.DeleteBook(isbn) // Use the repository to get the books
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, nil), nil
}

// BooksIsbnGet - Get a specific book by ISBN
func (s *DefaultAPIService) BooksIsbnGet(ctx context.Context, isbn string) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	book, err := s.Repo.GetBookByISBN(isbn) // Use the repository to get the books
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, convertBookToAPIFormat(*book)), nil
}

// BooksIsbnPatch - Update a book by ISBN
func (s *DefaultAPIService) BooksIsbnPatch(ctx context.Context, isbn string, book models.Book) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	// Check if the provided ISBN in the request path matches the ISBN in the body
	if book.Isbn != isbn {
		return common.Response(http.StatusBadRequest, nil), errors.New("ISBN in the path does not match ISBN in the body")
	}

	// Call the repository method to update the book
	dbBook := convertToDBBook(book)
	err := s.Repo.UpdateBook(&dbBook)
	if err != nil {
		if err == sql.ErrNoRows {
			return common.Response(http.StatusNotFound, nil), errors.New("book not found")
		}
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, nil), nil
}

// BooksPost - Add a new book
func (s *DefaultAPIService) BooksPost(ctx context.Context, book models.Book) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	dbBook := convertToDBBook(book)

	err := s.Repo.CreateBook(&dbBook)
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to add book: %w", err)
	}

	return common.Response(http.StatusCreated, nil), nil
}

// Convert Book to db.Book
func convertToDBBook(book models.Book) db.Book {
	dbBook := db.Book{
		ISBN:            book.Isbn,
		Name:            book.Name,
		Tags:            book.Tags,
		AuthorName:      book.AuthorName,
		DateOfPublish:   book.DateOfPublish,
		PublishingHouse: book.PublishingHouse,
		NumberOfPages:   int(book.NumberOfPages), // Convert int32 to int
		Cost:            float64(book.Cost),      // Convert float32 to float64
	}
	return dbBook
}

// convertBookToAPIFormat converts internal book format to API format
func convertBookToAPIFormat(book db.Book) map[string]interface{} {
	// Custom date format
	dateFormat := "01/02/06" // Date format: MM/DD/YY
	parsedDate, _ := time.Parse("2006-01-02", book.DateOfPublish)
	formattedDate := parsedDate.Format(dateFormat)

	// Prepare the API response
	return map[string]interface{}{
		"isbn":             book.ISBN,
		"name":             book.Name,
		"tags":             book.Tags,
		"author_name":      book.AuthorName,
		"date_of_publish":  formattedDate,
		"publishing_house": book.PublishingHouse,
		"number_of_pages":  book.NumberOfPages,
		"cost":             book.Cost,
	}
}
