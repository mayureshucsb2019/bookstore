package models

type CustomerAddress struct {
	Unit string `json:"unit,omitempty"`

	StreetName string `json:"street_name,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Country string `json:"country,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	Landmark string `json:"landmark,omitempty"`
}

// AssertCustomerAddressRequired checks if the required fields are not zero-ed
func AssertCustomerAddressRequired(obj CustomerAddress) error {
	return nil
}

// AssertCustomerAddressConstraints checks if the values respects the defined constraints
func AssertCustomerAddressConstraints(obj CustomerAddress) error {
	return nil
}
