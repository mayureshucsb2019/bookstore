// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Bookstore API
 *
 * API for managing books in an online bookstore.
 *
 * API version: 1.0.0
 */

package openapi

import (
	"context"
	"net/http"
)



// DefaultAPIRouter defines the required methods for binding the api requests to a responses for the DefaultAPI
// The DefaultAPIRouter implementation should parse necessary information from the http request,
// pass the data to a DefaultAPIServicer to perform the required actions, then write the service results to the http response.
type DefaultAPIRouter interface { 
	BooksGet(http.ResponseWriter, *http.Request)
	BooksIsbnDelete(http.ResponseWriter, *http.Request)
	BooksIsbnGet(http.ResponseWriter, *http.Request)
	BooksIsbnPatch(http.ResponseWriter, *http.Request)
	BooksPost(http.ResponseWriter, *http.Request)
}


// DefaultAPIServicer defines the api actions for the DefaultAPI service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type DefaultAPIServicer interface { 
	BooksGet(context.Context, int32, int32) (ImplResponse, error)
	BooksIsbnDelete(context.Context, string) (ImplResponse, error)
	BooksIsbnGet(context.Context, string) (ImplResponse, error)
	BooksIsbnPatch(context.Context, string, Book) (ImplResponse, error)
	BooksPost(context.Context, Book) (ImplResponse, error)
}
