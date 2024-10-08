package models

type BooksGet200Response struct {

	// Total number of books available.
	TotalItems int32 `json:"totalItems,omitempty"`

	// Total number of pages.
	TotalPages int32 `json:"totalPages,omitempty"`

	// The current page number.
	CurrentPage int32 `json:"currentPage,omitempty"`

	// The number of items per page.
	PageSize int32 `json:"pageSize,omitempty"`

	Books []Book `json:"books,omitempty"`
}

// AssertBooksGet200ResponseRequired checks if the required fields are not zero-ed
func AssertBooksGet200ResponseRequired(obj BooksGet200Response) error {
	for _, el := range obj.Books {
		if err := AssertBookRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertBooksGet200ResponseConstraints checks if the values respects the defined constraints
func AssertBooksGet200ResponseConstraints(obj BooksGet200Response) error {
	for _, el := range obj.Books {
		if err := AssertBookConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
