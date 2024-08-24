package models

import "github.com/mayureshucsb2019/bookstore/go/common"

type Author struct {
	Id string `json:"id"`

	Name AuthorName `json:"name"`

	Dob string `json:"dob"`

	Address AuthorAddress `json:"address"`

	Languages []string `json:"languages"`
}

// AssertAuthorRequired checks if the required fields are not zero-ed
func AssertAuthorRequired(obj Author) error {
	elements := map[string]interface{}{
		"id":        obj.Id,
		"name":      obj.Name,
		"dob":       obj.Dob,
		"address":   obj.Address,
		"languages": obj.Languages,
	}
	for name, el := range elements {
		if isZero := common.IsZeroValue(el); isZero {
			return &common.RequiredError{Field: name}
		}
	}

	if err := AssertAuthorNameRequired(obj.Name); err != nil {
		return err
	}
	if err := AssertAuthorAddressRequired(obj.Address); err != nil {
		return err
	}
	return nil
}

// AssertAuthorConstraints checks if the values respects the defined constraints
func AssertAuthorConstraints(obj Author) error {
	if err := AssertAuthorNameConstraints(obj.Name); err != nil {
		return err
	}
	if err := AssertAuthorAddressConstraints(obj.Address); err != nil {
		return err
	}
	return nil
}
