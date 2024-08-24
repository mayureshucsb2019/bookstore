package models

type AuthorName struct {
	FirstName string `json:"first_name,omitempty"`

	MiddleName string `json:"middle_name,omitempty"`

	LastName string `json:"last_name,omitempty"`
}

// AssertAuthorNameRequired checks if the required fields are not zero-ed
func AssertAuthorNameRequired(obj AuthorName) error {
	return nil
}

// AssertAuthorNameConstraints checks if the values respects the defined constraints
func AssertAuthorNameConstraints(obj AuthorName) error {
	return nil
}
