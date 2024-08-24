package service

import (
	"context"
	"net/http"

	"github.com/mayureshucsb2019/bookstore/service/author/models"
	"github.com/mayureshucsb2019/bookstore/service/common"
)

// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface {
	AuthorsGet(http.ResponseWriter, *http.Request)
	AuthorsIdDelete(http.ResponseWriter, *http.Request)
	AuthorsIdGet(http.ResponseWriter, *http.Request)
	AuthorsIdPatch(http.ResponseWriter, *http.Request)
	AuthorsPost(http.ResponseWriter, *http.Request)
}

// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface {
	AuthorsGet(context.Context, int32, int32) (common.ImplResponse, error)
	AuthorsIdDelete(context.Context, string) (common.ImplResponse, error)
	AuthorsIdGet(context.Context, string) (common.ImplResponse, error)
	AuthorsIdPatch(context.Context, string, models.Author) (common.ImplResponse, error)
	AuthorsPost(context.Context, models.Author) (common.ImplResponse, error)
}
