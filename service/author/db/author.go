package db

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

// Author struct represents the structure of a Author record in the database.
type Author struct {
	ID         string         `json:"id" db:"id"`
	FirstName  string         `json:"first_name" db:"first_name"`
	MiddleName sql.NullString `json:"middle_name" db:"middle_name"` // Nullable string
	LastName   string         `json:"last_name" db:"last_name"`
	DOB        string         `json:"dob" db:"dob"`   // Date in string format
	Unit       sql.NullString `json:"unit" db:"unit"` // Nullable string
	StreetName sql.NullString `json:"street_name" db:"street_name"`
	City       sql.NullString `json:"city" db:"city"`
	State      sql.NullString `json:"state" db:"state"`
	Country    sql.NullString `json:"country" db:"country"`
	Zipcode    sql.NullString `json:"zipcode" db:"zipcode"`
	Landmark   sql.NullString `json:"landmark" db:"landmark"`
	Languages  []string       `json:"languages" db:"languages"` // JSON array of strings
}

// Scan method to handle the JSON decoding for the Languages field
func (a *Author) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		return json.Unmarshal([]byte(src.(string)), &a.Languages)
	case []byte:
		return json.Unmarshal(src.([]byte), &a.Languages)
	default:
		return nil
	}
}

// AuthorRepository provides access to the Author storage.
type AuthorRepository struct {
	DB *sql.DB
}

// NewAuthorRepository creates a new AuthorRepository with a database connection.
func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

// CreateAuthor inserts a new Author into the database.
func (r *AuthorRepository) CreateAuthor(Author *Author) error {
	return nil
}

// GetAuthorByISBN retrieves a Author from the database by its ISBN.
func (r *AuthorRepository) GetAuthorByID(id string) (*Author, error) {
	return &Author{}, nil
}

// UpdateAuthor updates an existing Author record in the database.
func (r *AuthorRepository) UpdateAuthor(Author *Author) error {
	return nil
}

// DeleteAuthor removes a Author from the database by its ISBN.
func (r *AuthorRepository) DeleteAuthor(id string) error {
	return nil
}

// GetAllAuthors retrieves all Authors from the database.
func (r *AuthorRepository) GetAllAuthors() ([]Author, error) {
	return []Author{}, nil
}
