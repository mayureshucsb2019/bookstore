package models

type AuthorAddress struct {
	Unit string `json:"unit,omitempty"`

	StreetName string `json:"street_name,omitempty"`

	City string `json:"city,omitempty"`

	State string `json:"state,omitempty"`

	Country string `json:"country,omitempty"`

	Zipcode string `json:"zipcode,omitempty"`

	Landmark string `json:"landmark,omitempty"`
}

// AssertAuthorAddressRequired checks if the required fields are not zero-ed
func AssertAuthorAddressRequired(obj AuthorAddress) error {
	return nil
}

// AssertAuthorAddressConstraints checks if the values respects the defined constraints
func AssertAuthorAddressConstraints(obj AuthorAddress) error {
	return nil
}
