package models

import "github.com/mayureshucsb2019/bookstore/go/common"

type Book struct {
	Isbn string `json:"isbn"`

	Name string `json:"name"`

	Tags []string `json:"tags,omitempty"`

	AuthorName string `json:"author_name"`

	DateOfPublish string `json:"date_of_publish"`

	PublishingHouse string `json:"publishing_house,omitempty"`

	NumberOfPages int32 `json:"number_of_pages,omitempty"`

	Cost float32 `json:"cost,omitempty"`
}

// AssertBookRequired checks if the required fields are not zero-ed
func AssertBookRequired(obj Book) error {
	elements := map[string]interface{}{
		"isbn":            obj.Isbn,
		"name":            obj.Name,
		"author_name":     obj.AuthorName,
		"date_of_publish": obj.DateOfPublish,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	return nil
}

// AssertBookConstraints checks if the values respects the defined constraints
func AssertBookConstraints(obj Book) error {
	return nil
}
