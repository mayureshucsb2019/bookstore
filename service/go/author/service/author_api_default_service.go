package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/mayureshucsb2019/bookstore/go/author/db"
	"github.com/mayureshucsb2019/bookstore/go/author/models"
	"github.com/mayureshucsb2019/bookstore/go/common"
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
	// TODO - update AuthorsGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AuthorsGet200Response{}) or use other options such as http.Ok ...
	// return Response(200, AuthorsGet200Response{}), nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("AuthorsGet method not implemented")
}

// AuthorsIdDelete - Delete an author by ID
func (s *DefaultAPIService) AuthorsIdDelete(ctx context.Context, id string) (common.ImplResponse, error) {
	// TODO - update AuthorsIdDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("AuthorsIdDelete method not implemented")
}

// AuthorsIdGet - Get a specific author by ID
func (s *DefaultAPIService) AuthorsIdGet(ctx context.Context, id string) (common.ImplResponse, error) {
	// TODO - update AuthorsIdGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Author{}) or use other options such as http.Ok ...
	// return Response(200, Author{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("AuthorsIdGet method not implemented")
}

// AuthorsIdPatch - Update an author by ID
func (s *DefaultAPIService) AuthorsIdPatch(ctx context.Context, id string, author models.Author) (common.ImplResponse, error) {
	// TODO - update AuthorsIdPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("AuthorsIdPatch method not implemented")
}

// AuthorsPost - Add a new author
func (s *DefaultAPIService) AuthorsPost(ctx context.Context, author models.Author) (common.ImplResponse, error) {
	// TODO - update AuthorsPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return Response(201, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("AuthorsPost method not implemented")
}
