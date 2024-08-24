package openapi

import (
	"context"
	"net/http"

	"github.com/mayureshucsb2019/bookstore/service/common"
	"github.com/mayureshucsb2019/bookstore/service/customer/models"
)

// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface {
	CustomersEmailDelete(http.ResponseWriter, *http.Request)
	CustomersEmailGet(http.ResponseWriter, *http.Request)
	CustomersEmailPatch(http.ResponseWriter, *http.Request)
	CustomersGet(http.ResponseWriter, *http.Request)
	CustomersPost(http.ResponseWriter, *http.Request)
}

// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface {
	CustomersEmailDelete(context.Context, string) (common.ImplResponse, error)
	CustomersEmailGet(context.Context, string) (common.ImplResponse, error)
	CustomersEmailPatch(context.Context, string, models.Customer) (common.ImplResponse, error)
	CustomersGet(context.Context, int32, int32) (common.ImplResponse, error)
	CustomersPost(context.Context, models.Customer) (common.ImplResponse, error)
}
