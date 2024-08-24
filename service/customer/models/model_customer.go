package models

import (
	"time"

	"github.com/mayureshucsb2019/bookstore/service/common"
)

type Customer struct {
	Email string `json:"email"`

	FirstName string `json:"first_name"`

	MiddleName string `json:"middle_name,omitempty"`

	LastName string `json:"last_name"`

	PhoneNumber string `json:"phone_number,omitempty"`

	Dob string `json:"dob"`

	UnitNo string `json:"unit_no,omitempty"`

	StreetName string `json:"street_name,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Country string `json:"country,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	Landmark string `json:"landmark,omitempty"`

	RegistrationDate time.Time `json:"registration_date,omitempty"`

	LastLogin time.Time `json:"last_login,omitempty"`

	Status string `json:"status,omitempty"`

	Notes string `json:"notes,omitempty"`
}

// AssertCustomerRequired checks if the required fields are not zero-ed
func AssertCustomerRequired(obj Customer) error {
	elements := map[string]interface{}{
		"email":      obj.Email,
		"first_name": obj.FirstName,
		"last_name":  obj.LastName,
		"dob":        obj.Dob,
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
