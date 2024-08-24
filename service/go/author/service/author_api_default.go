package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mayureshucsb2019/bookstore/service/go/author/models"
	"github.com/mayureshucsb2019/bookstore/service/go/common"
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
		"AuthorsGet": common.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/authors",
			HandlerFunc: c.AuthorsGet,
		},
		"AuthorsIdDelete": common.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "/authors/{id}",
			HandlerFunc: c.AuthorsIdDelete,
		},
		"AuthorsIdGet": common.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/authors/{id}",
			HandlerFunc: c.AuthorsIdGet,
		},
		"AuthorsIdPatch": common.Route{
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/authors/{id}",
			HandlerFunc: c.AuthorsIdPatch,
		},
		"AuthorsPost": common.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/authors",
			HandlerFunc: c.AuthorsPost,
		},
	}
}

// AuthorsGet - Get a list of authors
func (c *DefaultAPIController) AuthorsGet(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.AuthorsGet(r.Context(), pageNumberParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AuthorsIdDelete - Delete an author by ID
func (c *DefaultAPIController) AuthorsIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "id"}, nil)
		return
	}
	result, err := c.service.AuthorsIdDelete(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AuthorsIdGet - Get a specific author by ID
func (c *DefaultAPIController) AuthorsIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "id"}, nil)
		return
	}
	result, err := c.service.AuthorsIdGet(r.Context(), idParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AuthorsIdPatch - Update an author by ID
func (c *DefaultAPIController) AuthorsIdPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idParam := params["id"]
	if idParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "id"}, nil)
		return
	}
	authorParam := models.Author{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&authorParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertAuthorRequired(authorParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertAuthorConstraints(authorParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AuthorsIdPatch(r.Context(), idParam, authorParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// AuthorsPost - Add a new author
func (c *DefaultAPIController) AuthorsPost(w http.ResponseWriter, r *http.Request) {
	authorParam := models.Author{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&authorParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertAuthorRequired(authorParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertAuthorConstraints(authorParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AuthorsPost(r.Context(), authorParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}
