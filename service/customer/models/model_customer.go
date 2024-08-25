package models

import "github.com/mayureshucsb2019/bookstore/service/common"

type Customer struct {
	Email string `json:"email"`

	Name CustomerName `json:"name"`

	PhoneNumber string `json:"phone_number,omitempty"`

	DOB string `json:"dob"`

	Address CustomerAddress `json:"address"`

	Status string `json:"status,omitempty"`

	Notes string `json:"notes,omitempty"`

	Languages []string `json:"languages"`
}

// AssertCustomerRequired checks if the required fields are not zero-ed
func AssertCustomerRequired(obj Customer) error {
	elements := map[string]interface{}{
		"email":      obj.Email,
		"first_name": obj.Name.FirstName,
		"last_name":  obj.Name.LastName,
		"dob":        obj.DOB,
		"languages":  obj.Languages,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCustomerConstraints checks if the values respects the defined constraints
func AssertCustomerConstraints(obj Customer) error {
	return nil
}
