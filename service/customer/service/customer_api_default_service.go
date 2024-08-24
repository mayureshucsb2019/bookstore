package openapi

import (
	"context"
	"errors"
	"net/http"

	"github.com/mayureshucsb2019/bookstore/service/common"
	"github.com/mayureshucsb2019/bookstore/service/customer/db"
	"github.com/mayureshucsb2019/bookstore/service/customer/models"
)

// DefaultAPIService is a service that implements the logic for the DefaultAPIServicer
// This service should implement the business logic for every endpoint for the DefaultAPI API.
// Include any external packages or services that will be required by this service.
type DefaultAPIService struct {
	Repo *db.CustomerRepository // Add a field to hold the repository
}

// NewDefaultAPIService creates a default api service
func NewDefaultAPIService(repo *db.CustomerRepository) *DefaultAPIService {
	return &DefaultAPIService{
		Repo: repo,
	}
}

// CustomersEmailDelete - Delete a customer by email
func (s *DefaultAPIService) CustomersEmailDelete(ctx context.Context, email string) (common.ImplResponse, error) {
	// TODO - update CustomersEmailDelete with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("CustomersEmailDelete method not implemented")
}

// CustomersEmailGet - Get a specific customer by email
func (s *DefaultAPIService) CustomersEmailGet(ctx context.Context, email string) (common.ImplResponse, error) {
	// TODO - update CustomersEmailGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Customer{}) or use other options such as http.Ok ...
	// return Response(200, Customer{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("CustomersEmailGet method not implemented")
}

// CustomersEmailPatch - Update a customer by email
func (s *DefaultAPIService) CustomersEmailPatch(ctx context.Context, email string, customer models.Customer) (common.ImplResponse, error) {
	// TODO - update CustomersEmailPatch with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("CustomersEmailPatch method not implemented")
}

// CustomersGet - Get a paginated list of customers
func (s *DefaultAPIService) CustomersGet(ctx context.Context, pageNumber int32, pageSize int32) (common.ImplResponse, error) {
	// TODO - update CustomersGet with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, CustomersGet200Response{}) or use other options such as http.Ok ...
	// return Response(200, CustomersGet200Response{}), nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("CustomersGet method not implemented")
}

// CustomersPost - Add a new customer
func (s *DefaultAPIService) CustomersPost(ctx context.Context, customer models.Customer) (common.ImplResponse, error) {
	// TODO - update CustomersPost with the required logic for this service method.
	// Add api_default_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return Response(201, nil),nil

	return common.Response(http.StatusNotImplemented, nil), errors.New("CustomersPost method not implemented")
}
