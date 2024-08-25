package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Customer represents the structure of a Customer record in the database.
type Customer struct {
	Email            string         `json:"email" db:"email"`
	FirstName        string         `json:"first_name" db:"first_name"`
	MiddleName       sql.NullString `json:"middle_name" db:"middle_name"` // Nullable string
	LastName         string         `json:"last_name" db:"last_name"`
	PhoneNumber      sql.NullString `json:"phone_number" db:"phone_number"`
	Dob              string         `json:"dob" db:"dob"`         // Date in string format
	UnitNo           sql.NullString `json:"unit_no" db:"unit_no"` // Nullable string
	StreetName       sql.NullString `json:"street_name" db:"street_name"`
	City             sql.NullString `json:"city" db:"city"`
	State            sql.NullString `json:"state" db:"state"`
	Country          sql.NullString `json:"country" db:"country"`
	Zipcode          sql.NullString `json:"zipcode" db:"zipcode"`
	Landmark         sql.NullString `json:"landmark" db:"landmark"`
	RegistrationDate string         `json:"registration_date" db:"registration_date"` // Timestamp as string
	LastLogin        sql.NullString `json:"last_login" db:"last_login"`               // Timestamp as string
	Status           string         `json:"status" db:"status"`                       // ENUM value
	Notes            sql.NullString `json:"notes" db:"notes"`
	Languages        []string       `json:"languages" db:"languages"` // JSON array of strings
}

// Scan method to handle the JSON decoding for the Languages field
func (a *Customer) Scan(src interface{}) error {
	switch src.(type) {
	case string:
		return json.Unmarshal([]byte(src.(string)), &a.Languages)
	case []byte:
		return json.Unmarshal(src.([]byte), &a.Languages)
	default:
		return nil
	}
}

// CustomerRepository provides access to the Customer storage.
type CustomerRepository struct {
	DB *sql.DB
}

// NewCustomerRepository creates a new CustomerRepository with a database connection.
func NewCustomerRepository(db *sql.DB) *CustomerRepository {
	return &CustomerRepository{DB: db}
}

// CreateCustomer inserts a new Customer into the database.
func (r *CustomerRepository) CreateCustomer(customer *Customer) error {
	languagesJSON, err := json.Marshal(customer.Languages)
	if err != nil {
		log.Fatalf("Failed to marshal languages to JSON: %v", err)
	}
	// Prepare the SQL insert statement
	query := `
		INSERT INTO Customer (
			email, first_name, middle_name, last_name, phone_number, dob,
			unit_no, street_name, city, state, country, zipcode, landmark,
			last_login, status, notes, languages
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`

	// Execute the SQL statement
	_, err = r.DB.Exec(query,
		customer.Email,
		customer.FirstName,
		customer.MiddleName,
		customer.LastName,
		customer.PhoneNumber,
		customer.Dob,
		customer.UnitNo,
		customer.StreetName,
		customer.City,
		customer.State,
		customer.Country,
		customer.Zipcode,
		customer.Landmark,
		customer.LastLogin,
		customer.Status,
		customer.Notes,
		languagesJSON,
	)
	if err != nil {
		return fmt.Errorf("failed to insert customer: %w", err)
	}

	return nil
}

// GetCustomerByID retrieves a Customer from the database by its email.
func (r *CustomerRepository) GetCustomerByID(email string) (*Customer, error) {
	// Prepare the SQL select statement
	query := `SELECT * FROM Customer WHERE email = ?`

	// Declare a variable to hold the Customer data
	var customer Customer

	// Temporary variable to hold the JSON data
	var languagesJSON []byte

	// Execute the query and scan the result into the customer variable
	row := r.DB.QueryRow(query, email)

	err := row.Scan(
		&customer.Email,
		&customer.FirstName,
		&customer.MiddleName,
		&customer.LastName,
		&customer.PhoneNumber,
		&customer.Dob,
		&customer.UnitNo,
		&customer.StreetName,
		&customer.City,
		&customer.State,
		&customer.Country,
		&customer.Zipcode,
		&customer.Landmark,
		&customer.RegistrationDate,
		&customer.LastLogin,
		&customer.Status,
		&customer.Notes,
		&languagesJSON,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// No customer found with the given email
			return nil, fmt.Errorf("customer with id %s not found", email)
		}
		// Other errors
		return nil, fmt.Errorf("failed to retrieve customer: %w", err)
	}

	// Unmarshal the JSON data into a []string slice
	var languages []string
	if len(languagesJSON) > 0 { // Check if languagesJSON is not empty
		if err := json.Unmarshal(languagesJSON, &languages); err != nil {
			return nil, fmt.Errorf("failed to unmarshal languages: %w", err)
		}
	}
	customer.Languages = languages

	return &customer, nil
}

// UpdateCustomer updates an existing Customer record in the database.
func (r *CustomerRepository) UpdateCustomer(customer *Customer) error {
	// Prepare the SQL update statement
	query := `
		UPDATE Customer
		SET
			first_name = ?, 
			middle_name = ?, 
			last_name = ?, 
			phone_number = ?, 
			dob = ?, 
			unit_no = ?, 
			street_name = ?, 
			city = ?, 
			state = ?, 
			country = ?, 
			zipcode = ?, 
			landmark = ?, 
			registration_date = ?, 
			last_login = ?, 
			status = ?, 
			notes = ?, 
			languages = ?
		WHERE email = ?
	`

	// Execute the SQL statement
	_, err := r.DB.Exec(query,
		customer.FirstName,
		customer.MiddleName,
		customer.LastName,
		customer.PhoneNumber,
		customer.Dob,
		customer.UnitNo,
		customer.StreetName,
		customer.City,
		customer.State,
		customer.Country,
		customer.Zipcode,
		customer.Landmark,
		customer.RegistrationDate,
		customer.LastLogin,
		customer.Status,
		customer.Notes,
		customer.Languages,
		customer.Email, // Email is used as the unique identifier
	)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}

	return nil
}

// DeleteCustomer removes a Customer from the database by their email.
func (r *CustomerRepository) DeleteCustomer(email string) error {
	// Prepare the SQL delete statement
	query := `DELETE FROM Customer WHERE email = ?`

	// Execute the SQL statement
	result, err := r.DB.Exec(query, email)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}

	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		// No customer found with the given email
		return nil
	}

	return nil
}

// GetAllCustomers retrieves all Customers from the database.
func (r *CustomerRepository) GetAllCustomers() ([]Customer, error) {
	// Prepare the SQL select statement
	query := `SELECT * FROM Customer`

	// Execute the query
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query customers: %w", err)
	}
	defer rows.Close()

	// Slice to hold all customers
	var customers []Customer

	// Iterate over the rows
	for rows.Next() {
		var customer Customer
		// Temporary variable to hold the JSON data
		var languagesJSON []byte

		err := rows.Scan(
			&customer.Email,
			&customer.FirstName,
			&customer.MiddleName,
			&customer.LastName,
			&customer.PhoneNumber,
			&customer.Dob,
			&customer.UnitNo,
			&customer.StreetName,
			&customer.City,
			&customer.State,
			&customer.Country,
			&customer.Zipcode,
			&customer.Landmark,
			&customer.RegistrationDate,
			&customer.LastLogin,
			&customer.Status,
			&customer.Notes,
			&languagesJSON,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer: %w", err)
		}

		// Unmarshal the languages JSON byte slice into a []string slice
		if len(languagesJSON) > 0 {
			if err := json.Unmarshal(languagesJSON, &customer.Languages); err != nil {
				return nil, fmt.Errorf("failed to unmarshal languages: %w", err)
			}
		}

		customers = append(customers, customer)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return customers, nil
}
