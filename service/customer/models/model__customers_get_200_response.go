package models

type CustomersGet200Response struct {

	// Total number of customers available.
	TotalItems int32 `json:"totalItems,omitempty"`

	// Total number of pages.
	TotalPages int32 `json:"totalPages,omitempty"`

	// The current page number.
	CurrentPage int32 `json:"currentPage,omitempty"`

	// The number of items per page.
	PageSize int32 `json:"pageSize,omitempty"`

	Customers []Customer `json:"customers,omitempty"`
}

// AssertCustomersGet200ResponseRequired checks if the required fields are not zero-ed
func AssertCustomersGet200ResponseRequired(obj CustomersGet200Response) error {
	for _, el := range obj.Customers {
		if err := AssertCustomerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCustomersGet200ResponseConstraints checks if the values respects the defined constraints
func AssertCustomersGet200ResponseConstraints(obj CustomersGet200Response) error {
	for _, el := range obj.Customers {
		if err := AssertCustomerConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
