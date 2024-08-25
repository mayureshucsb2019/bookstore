package openapi

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	err := s.Repo.DeleteCustomer(email)
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, nil), nil
}

// CustomersEmailGet - Get a specific customer by email
func (s *DefaultAPIService) CustomersEmailGet(ctx context.Context, email string) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	customer, err := s.Repo.GetCustomerByID(email)
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, convertDBToAPIResponse(*customer)), nil
}

// CustomersEmailPatch - Update a customer by email
func (s *DefaultAPIService) CustomersEmailPatch(ctx context.Context, email string, customer models.Customer) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	// Check if the provided ISBN in the request path matches the ISBN in the body
	if customer.Email != email {
		return common.Response(http.StatusBadRequest, nil), errors.New("email in the path does not match email in the body")
	}

	// Call the repository method to update the customer
	dbCustomer := convertApiToDBCustomer(customer)
	err := s.Repo.UpdateCustomer(&dbCustomer)
	if err != nil {
		if err == sql.ErrNoRows {
			return common.Response(http.StatusNotFound, nil), errors.New("customer not found")
		}
		return common.Response(http.StatusInternalServerError, nil), err
	}

	return common.Response(http.StatusOK, nil), nil
}

// CustomersGet - Get a paginated list of customers
func (s *DefaultAPIService) CustomersGet(ctx context.Context, pageNumber int32, pageSize int32) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	customers, err := s.Repo.GetAllCustomers()
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), err
	}
	var customerResp []models.Customer
	for _, customer := range customers {
		customerResp = append(customerResp, convertDBToAPIResponse(customer))
	}

	return common.Response(http.StatusOK, customerResp), nil
}

// CustomersPost - Add a new customer
func (s *DefaultAPIService) CustomersPost(ctx context.Context, customer models.Customer) (common.ImplResponse, error) {
	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil
	dbCustomer := convertApiToDBCustomer(customer)

	err := s.Repo.CreateCustomer(&dbCustomer)
	if err != nil {
		return common.Response(http.StatusInternalServerError, nil), fmt.Errorf("failed to add customer: %w", err)
	}

	return common.Response(http.StatusCreated, nil), nil
}

// ConvertApiToDBCustomer converts an API Customer struct to a database Customer struct.
func convertApiToDBCustomer(apiCustomer models.Customer) db.Customer {
	return db.Customer{
		Email:       apiCustomer.Email,
		FirstName:   apiCustomer.Name.FirstName,
		MiddleName:  common.NullStringOrNil(apiCustomer.Name.MiddleName),
		LastName:    apiCustomer.Name.LastName,
		PhoneNumber: common.NullStringOrNil(apiCustomer.PhoneNumber),
		Dob:         apiCustomer.DOB,
		UnitNo:      common.NullStringOrNil(apiCustomer.Address.Unit),
		StreetName:  common.NullStringOrNil(apiCustomer.Address.StreetName),
		City:        common.NullStringOrNil(apiCustomer.Address.City),
		State:       common.NullStringOrNil(apiCustomer.Address.State),
		Country:     common.NullStringOrNil(apiCustomer.Address.Country),
		Zipcode:     common.NullStringOrNil(apiCustomer.Address.Zipcode),
		Landmark:    common.NullStringOrNil(apiCustomer.Address.Landmark),
		Status:      apiCustomer.Status,
		Notes:       common.NullStringOrNil(apiCustomer.Notes),
		Languages:   apiCustomer.Languages,
	}
}

// convertDBToAPIResponse converts a DBCustomer struct to an APICustomer struct.
func convertDBToAPIResponse(dbCustomer db.Customer) models.Customer {
	return models.Customer{
		Email: dbCustomer.Email,
		Name: models.CustomerName{
			FirstName:  dbCustomer.FirstName,
			MiddleName: common.StringOrEmpty(dbCustomer.MiddleName),
			LastName:   dbCustomer.LastName,
		},
		PhoneNumber: common.StringOrEmpty(dbCustomer.PhoneNumber),
		DOB:         dbCustomer.Dob,
		Address: models.CustomerAddress{
			Unit:       common.StringOrEmpty(dbCustomer.UnitNo),
			StreetName: common.StringOrEmpty(dbCustomer.StreetName),
			City:       common.StringOrEmpty(dbCustomer.City),
			State:      common.StringOrEmpty(dbCustomer.State),
			Country:    common.StringOrEmpty(dbCustomer.Country),
			Zipcode:    common.StringOrEmpty(dbCustomer.Zipcode),
			Landmark:   common.StringOrEmpty(dbCustomer.Landmark),
		},
		Status:    dbCustomer.Status,
		Notes:     common.StringOrEmpty(dbCustomer.Notes),
		Languages: dbCustomer.Languages,
	}
}
