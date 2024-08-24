package models

type CustomerName struct {
	FirstName string `json:"first_name,omitempty"`

	MiddleName string `json:"middle_name,omitempty"`

	LastName string `json:"last_name,omitempty"`
}

// AssertCustomerNameRequired checks if the required fields are not zero-ed
func AssertCustomerNameRequired(obj CustomerName) error {
	return nil
}

// AssertCustomerNameConstraints checks if the values respects the defined constraints
func AssertCustomerNameConstraints(obj CustomerName) error {
	return nil
}
