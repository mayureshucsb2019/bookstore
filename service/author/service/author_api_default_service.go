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
		MiddleName: common.NullStringOrNil(author.Name.MiddleName),
		LastName:   author.Name.LastName,
		DOB:        common.NullStringOrNil(author.DOB),
		UnitNo:     common.NullStringOrNil(author.Address.Unit),
		StreetName: common.NullStringOrNil(author.Address.StreetName),
		City:       common.NullStringOrNil(author.Address.City),
		State:      common.NullStringOrNil(author.Address.State),
		Country:    common.NullStringOrNil(author.Address.Country),
		Zipcode:    common.NullStringOrNil(author.Address.Zipcode),
		Landmark:   common.NullStringOrNil(author.Address.Landmark),
		Languages:  author.Languages,
	}
}

// ConvertDBToAPIResponse converts the DB model to the API model
func convertDBToAPIResponse(db db.Author) models.Author {
	return models.Author{
		Id: db.ID,
		Name: models.AuthorName{
			FirstName:  db.FirstName,
			MiddleName: common.StringOrEmpty(db.MiddleName),
			LastName:   db.LastName,
		},
		DOB: common.StringOrEmpty(db.DOB),
		Address: models.AuthorAddress{
			Unit:       common.StringOrEmpty(db.UnitNo),
			StreetName: common.StringOrEmpty(db.StreetName),
			City:       common.StringOrEmpty(db.City),
			State:      common.StringOrEmpty(db.State),
			Country:    common.StringOrEmpty(db.Country),
			Zipcode:    common.StringOrEmpty(db.Zipcode),
			Landmark:   common.StringOrEmpty(db.Landmark),
		},
		Languages: db.Languages,
	}
}
