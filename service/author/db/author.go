package db

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Author represents the structure of an Author record in the database.
type Author struct {
	ID         string         `json:"id" db:"id"`
	FirstName  string         `json:"first_name" db:"first_name"`
	MiddleName sql.NullString `json:"middle_name" db:"middle_name"`
	LastName   string         `json:"last_name" db:"last_name"`
	DOB        string         `json:"dob" db:"dob"`
	Unit       sql.NullString `json:"unit" db:"unit"`
	StreetName sql.NullString `json:"street_name" db:"street_name"`
	City       sql.NullString `json:"city" db:"city"`
	State      sql.NullString `json:"state" db:"state"`
	Country    sql.NullString `json:"country" db:"country"`
	Zipcode    sql.NullString `json:"zipcode" db:"zipcode"`
	Landmark   sql.NullString `json:"landmark" db:"landmark"`
	Languages  []string       `json:"languages" db:"-"` // Use []string for JSON unmarshalling
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
func (r *AuthorRepository) CreateAuthor(author *Author) error {
	// Prepare the SQL query for inserting a new Author
	query := `
		INSERT INTO Authors (
			id, first_name, middle_name, last_name, dob, unit_no, 
			street_name, city, state, country, zipcode, landmark, languages
		) VALUES (
			?, ?, ?, ?, ?, ?, 
			?, ?, ?, ?, ?, ?, ?
		)
	`
	// Convert the languages slice to a JSON string
	languagesJSON, err := json.Marshal(author.Languages)
	if err != nil {
		return fmt.Errorf("failed to marshal languages: %w", err)
	}

	// Execute the query with the provided author data
	_, err = r.DB.Exec(query,
		author.ID,
		author.FirstName,
		author.MiddleName,
		author.LastName,
		author.DOB,
		author.Unit,
		author.StreetName,
		author.City,
		author.State,
		author.Country,
		author.Zipcode,
		author.Landmark,
		languagesJSON,
	)

	return err
}

// GetAuthorByID retrieves an Author by its ID from the database.
func (r *AuthorRepository) GetAuthorByID(id string) (*Author, error) {
	// Prepare the SQL query for selecting an Author by ID
	query := `SELECT * FROM Authors WHERE id = ?`

	// Declare a variable to hold the Author data
	var author Author

	// Temporary variable to hold the JSON data
	var languagesJSON []byte

	// Execute the query and scan the result into the author variable
	row := r.DB.QueryRow(query, id)

	err := row.Scan(
		&author.ID,
		&author.FirstName,
		&author.MiddleName,
		&author.LastName,
		&author.DOB,
		&author.Unit,
		&author.StreetName,
		&author.City,
		&author.State,
		&author.Country,
		&author.Zipcode,
		&author.Landmark,
		&languagesJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("author with id %s not found", id)
		}
		return nil, fmt.Errorf("failed to get author by id: %w", err)
	}

	// Unmarshal the JSON data into a []string slice
	var languages []string
	if len(languagesJSON) > 0 { // Check if languagesJSON is not empty
		if err := json.Unmarshal(languagesJSON, &languages); err != nil {
			return nil, fmt.Errorf("failed to unmarshal languages: %w", err)
		}
	}
	author.Languages = languages

	return &author, nil
}

// UpdateAuthor updates an existing Author record in the database.
func (r *AuthorRepository) UpdateAuthor(author *Author) error {
	// Prepare the SQL query for updating an Author record
	query := `
		UPDATE Authors
		SET 
			first_name = ?, 
			middle_name = ?, 
			last_name = ?, 
			dob = ?, 
			unit_no = ?, 
			street_name = ?, 
			city = ?, 
			state = ?, 
			country = ?, 
			zipcode = ?, 
			landmark = ?, 
			languages = ?
		WHERE 
			id = ?
	`

	// Marshal the Languages slice to JSON
	languagesJSON, err := json.Marshal(author.Languages)
	if err != nil {
		return fmt.Errorf("failed to marshal languages: %w", err)
	}

	// Execute the query
	result, err := r.DB.Exec(
		query,
		author.FirstName,
		author.MiddleName,
		author.LastName,
		author.DOB,
		author.Unit,
		author.StreetName,
		author.City,
		author.State,
		author.Country,
		author.Zipcode,
		author.Landmark,
		languagesJSON,
		author.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update author: %w", err)
	}

	// Check if the update affected any rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no changes were made to author with id %s", author.ID)
	}

	return nil
}

// DeleteAuthor removes an Author from the database by its ID.
func (r *AuthorRepository) DeleteAuthor(id string) error {
	// Prepare the SQL query for deleting an Author record
	query := `DELETE FROM Authors WHERE id = ?`

	// Execute the query
	result, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete author: %w", err)
	}

	// Check if the delete operation affected any rows
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no author found with id %s", id)
	}

	return nil
}

// GetAllAuthors retrieves all Authors from the database.
func (r *AuthorRepository) GetAllAuthors() ([]Author, error) {
	// Prepare the SQL query for selecting all Author records
	query := `SELECT * FROM Authors`

	// Execute the query
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query authors: %w", err)
	}
	defer rows.Close()

	// Declare a slice to hold the authors
	var authors []Author

	// Iterate over the rows
	for rows.Next() {
		var author Author
		var languagesJSON []byte

		// Scan the row into the author struct
		if err := rows.Scan(
			&author.ID,
			&author.FirstName,
			&author.MiddleName,
			&author.LastName,
			&author.DOB,
			&author.Unit,
			&author.StreetName,
			&author.City,
			&author.State,
			&author.Country,
			&author.Zipcode,
			&author.Landmark,
			&languagesJSON,
		); err != nil {
			return nil, fmt.Errorf("failed to scan author: %w", err)
		}

		// Unmarshal the languages JSON byte slice into a []string slice
		if len(languagesJSON) > 0 {
			if err := json.Unmarshal(languagesJSON, &author.Languages); err != nil {
				return nil, fmt.Errorf("failed to unmarshal languages: %w", err)
			}
		}

		// Add the author to the slice
		authors = append(authors, author)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return authors, nil
}
