package models

type AuthorsGet200Response struct {

	// Total number of authors available.
	TotalItems int32 `json:"totalItems,omitempty"`

	// Total number of pages.
	TotalPages int32 `json:"totalPages,omitempty"`

	// The current page number.
	CurrentPage int32 `json:"currentPage,omitempty"`

	// The number of items per page.
	PageSize int32 `json:"pageSize,omitempty"`

	Authors []Author `json:"authors,omitempty"`
}

// AssertAuthorsGet200ResponseRequired checks if the required fields are not zero-ed
func AssertAuthorsGet200ResponseRequired(obj AuthorsGet200Response) error {
	for _, el := range obj.Authors {
		if err := AssertAuthorRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAuthorsGet200ResponseConstraints checks if the values respects the defined constraints
func AssertAuthorsGet200ResponseConstraints(obj AuthorsGet200Response) error {
	for _, el := range obj.Authors {
		if err := AssertAuthorConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
