package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/mayureshucsb2019/bookstore/service/author/db"
	"github.com/mayureshucsb2019/bookstore/service/author/models"
	"github.com/mayureshucsb2019/bookstore/service/common"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
	Repo *db.AuthorRepository // Add a field to hold the repository
}

// NewDefaultAPIService creates a default API service with the given repository.
func NewDefaultAPIService(repo *db.AuthorRepository) *DefaultAPIService {
	return &DefaultAPIService{
		Repo: repo,
	}
}

// AuthorsGet - Get a list of authors
func (s *DefaultAPIService) AuthorsGet(ctx context.Context, pageNumber int32, pageSize int32) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	authors, err := s.Repo.GetAllAuthors() // Use the repository to get the books
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}
	var authorResp []models.Author
	for _, author := range authors {
		authorResp = append(authorResp, convertDBToAPIResponse(author))
	}

	return common.Response(http.StatusOK, authorResp), nil
}

// AuthorsIdDelete - Delete an author by ID
func (s *DefaultAPIService) AuthorsIdDelete(ctx context.Context, id string) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	err := s.Repo.DeleteAuthor(id) // Use the repository to get the books
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, nil), nil
}

// AuthorsIdGet - Get a specific author by ID
func (s *DefaultAPIService) AuthorsIdGet(ctx context.Context, id string) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	author, err := s.Repo.GetAuthorByID(id) // Use the repository to get the books
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, convertDBToAPIResponse(*author)), nil
}

// AuthorsIdPatch - Update an author by ID
func (s *DefaultAPIService) AuthorsIdPatch(ctx context.Context, id string, author models.Author) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	// Check if the provided ISBN in the request path matches the ISBN in the body
	if author.Id != id {
		return common.Response(http.StatusBadRequest, nil), errors.New("id in the path does not match id in the body")
	}

	// Call the repository method to update the book
	dbAuthor := convertApiToDBAuthor(author)
	err := s.Repo.UpdateAuthor(&dbAuthor)
	if err != nil {
		if err == sql.ErrNoRows {
			return common.Response(http.StatusNotFound, nil), errors.New("book not found")
		}
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, nil), nil
}

// AuthorsPost - Add a new author
func (s *DefaultAPIService) AuthorsPost(ctx context.Context, author models.Author) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	dbAuthor := convertApiToDBAuthor(author)

	err := s.Repo.CreateAuthor(&dbAuthor)
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to add author: %w", err)
	}

	return common.Response(http.StatusCreated, nil), nil
}

// ConvertToDBAuthor converts an API model Author to a database model Author.
func convertApiToDBAuthor(author models.Author) db.Author {
	return db.Author{
		ID:         author.Id,
		FirstName:  author.Name.FirstName,
		MiddleName: nullStringOrNil(author.Name.MiddleName),
		LastName:   author.Name.LastName,
		DOB:        nullStringOrNil(author.DOB),
		UnitNo:     nullStringOrNil(author.Address.Unit),
		StreetName: nullStringOrNil(author.Address.StreetName),
		City:       nullStringOrNil(author.Address.City),
		State:      nullStringOrNil(author.Address.State),
		Country:    nullStringOrNil(author.Address.Country),
		Zipcode:    nullStringOrNil(author.Address.Zipcode),
		Landmark:   nullStringOrNil(author.Address.Landmark),
		Languages:  author.Languages,
	}
}

// Convert a string to sql.NullString. Returns a NullString with valid value or null.
func nullStringOrNil(value string) sql.NullString {
	if value == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: value, Valid: true}
}

// ConvertDBToAPIResponse converts the DB model to the API model
func convertDBToAPIResponse(db db.Author) models.Author {
	return models.Author{
		Id: db.ID,
		Name: models.AuthorName{
			FirstName:  db.FirstName,
			MiddleName: stringOrEmpty(db.MiddleName),
			LastName:   db.LastName,
		},
		DOB: stringOrEmpty(db.DOB),
		Address: models.AuthorAddress{
			Unit:       stringOrEmpty(db.UnitNo),
			StreetName: stringOrEmpty(db.StreetName),
			City:       stringOrEmpty(db.City),
			State:      stringOrEmpty(db.State),
			Country:    stringOrEmpty(db.Country),
			Zipcode:    stringOrEmpty(db.Zipcode),
			Landmark:   stringOrEmpty(db.Landmark),
		},
		Languages: db.Languages,
	}
}

// Helper function to convert sql.NullString to string
func stringOrEmpty(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
