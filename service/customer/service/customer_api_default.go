package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mayureshucsb2019/bookstore/service/common"
	"github.com/mayureshucsb2019/bookstore/service/customer/models"
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
		"CustomersEmailDelete": common.Route{
			Method:      strings.ToUpper("Delete"),
			Pattern:     "/customers/{email}",
			HandlerFunc: c.CustomersEmailDelete,
		},
		"CustomersEmailGet": common.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/customers/{email}",
			HandlerFunc: c.CustomersEmailGet,
		},
		"CustomersEmailPatch": common.Route{
			Method:      strings.ToUpper("Patch"),
			Pattern:     "/customers/{email}",
			HandlerFunc: c.CustomersEmailPatch,
		},
		"CustomersGet": common.Route{
			Method:      strings.ToUpper("Get"),
			Pattern:     "/customers",
			HandlerFunc: c.CustomersGet,
		},
		"CustomersPost": common.Route{
			Method:      strings.ToUpper("Post"),
			Pattern:     "/customers",
			HandlerFunc: c.CustomersPost,
		},
	}
}

// CustomersEmailDelete - Delete a customer by email
func (c *DefaultAPIController) CustomersEmailDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailParam := params["email"]
	if emailParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "email"}, nil)
		return
	}
	result, err := c.service.CustomersEmailDelete(r.Context(), emailParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CustomersEmailGet - Get a specific customer by email
func (c *DefaultAPIController) CustomersEmailGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailParam := params["email"]
	if emailParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "email"}, nil)
		return
	}
	result, err := c.service.CustomersEmailGet(r.Context(), emailParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CustomersEmailPatch - Update a customer by email
func (c *DefaultAPIController) CustomersEmailPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	emailParam := params["email"]
	if emailParam == "" {
		c.errorHandler(w, r, &common.RequiredError{Field: "email"}, nil)
		return
	}
	customerParam := models.Customer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&customerParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertCustomerRequired(customerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertCustomerConstraints(customerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CustomersEmailPatch(r.Context(), emailParam, customerParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CustomersGet - Get a paginated list of customers
func (c *DefaultAPIController) CustomersGet(w http.ResponseWriter, r *http.Request) {
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
	result, err := c.service.CustomersGet(r.Context(), pageNumberParam, pageSizeParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}

// CustomersPost - Add a new customer
func (c *DefaultAPIController) CustomersPost(w http.ResponseWriter, r *http.Request) {
	customerParam := models.Customer{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&customerParam); err != nil {
		c.errorHandler(w, r, &common.ParsingError{Err: err}, nil)
		return
	}
	if err := models.AssertCustomerRequired(customerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	if err := models.AssertCustomerConstraints(customerParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.CustomersPost(r.Context(), customerParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	_ = common.EncodeJSONResponse(result.Body, &result.Code, w)
}
