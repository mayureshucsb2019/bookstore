package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mayureshucsb2019/bookstore/service/book/models"
	"github.com/mayureshucsb2019/bookstore/service/common"
)

// DefaultAPIController binds http requests to an api service and writes the service results to the http response
type DefaultAPIController struct {
	service      DefaultAPIServicer
	errorHandler common.ErrorHandler
}

// DefaultAPIOption for how the controller is set up.
type DefaultAPIOption func(*DefaultAPIController)

// WithDefaultAPIErrorHandler inject ErrorHandler into controller
func WithDefaultAPIErrorHandler(h common.ErrorHandler) DefaultAPIOption {
	return func(c *DefaultAPIController) {
		c.errorHandler = h
	}
}

// NewDefaultAPIController creates a default api controller
func NewDefaultAPIController(s DefaultAPIServicer, opts ...DefaultAPIOption) *DefaultAPIController {
	controller := &DefaultAPIController{
		service:      s,
		errorHandler: common.DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the DefaultAPIController
func (c *DefaultAPIController) Routes() common.Routes {
	return common.Routes{
		"BooksGet": common.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/books",
			HandlerFunc: c.BooksGet,
		},
		"BooksIsbnDelete": common.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "/books/{isbn}",
			HandlerFunc: c.BooksIsbnDelete,
		},
		"BooksIsbnGet": common.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/books/{isbn}",
			HandlerFunc: c.BooksIsbnGet,
		},
		"BooksIsbnPatch": common.Route{
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/books/{isbn}",
			HandlerFunc: c.BooksIsbnPatch,
		},
		"BooksPost": common.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/books",
			HandlerFunc: c.BooksPost,
		},
	}
}

// BooksGet - Get a paginated list of books
func (c *DefaultAPIController) BooksGet(w http.ResponseWriter, r *http.Request) {
	query, err := common.ParseQuery(r.URL.RawQuery)
	if err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	var pageNumberParam int32
	if query.Has("pageNumber") {
		param, err := common.ParseNumericParameter[int32](
			query.Get("pageNumber"),
			common.WithParse[int32](common.ParseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &common.ParsingError{Param: "pageNumber", Err: err}, nil)
			return
		}

		pageNumberParam = param
	} else {
		var param int32 = 1
		pageNumberParam = param
	}
	var pageSizeParam int32
	if query.Has("pageSize") {
		param, err := common.ParseNumericParameter[int32](
			query.Get("pageSize"),
			common.WithParse[int32](common.ParseInt32),
		)
		if err != nil {
			c.errorHandler(w, r, &common.ParsingError{Param: "pageSize", Err: err}, nil)
			return
		}

		pageSizeParam = param
	} else {
		var param int32 = 25
		pageSizeParam = param
	}
	result, err := c.service.BooksGet(r.Context(), pageNumberParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// BooksIsbnDelete - Delete a book by ISBN
func (c *DefaultAPIController) BooksIsbnDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	isbnParam := params["isbn"]
	if isbnParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "isbn"}, nil)
		return
	}
	result, err := c.service.BooksIsbnDelete(r.Context(), isbnParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// BooksIsbnGet - Get a specific book by ISBN
func (c *DefaultAPIController) BooksIsbnGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	isbnParam := params["isbn"]
	if isbnParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "isbn"}, nil)
		return
	}
	result, err := c.service.BooksIsbnGet(r.Context(), isbnParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// BooksIsbnPatch - Update a book by ISBN
func (c *DefaultAPIController) BooksIsbnPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	isbnParam := params["isbn"]
	if isbnParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "isbn"}, nil)
		return
	}
	bookParam := models.Book{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertBookRequired(bookParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertBookConstraints(bookParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.BooksIsbnPatch(r.Context(), isbnParam, bookParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// BooksPost - Add a new book
func (c *DefaultAPIController) BooksPost(w http.ResponseWriter, r *http.Request) {
	bookParam := models.Book{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&bookParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertBookRequired(bookParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertBookConstraints(bookParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.BooksPost(r.Context(), bookParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}
