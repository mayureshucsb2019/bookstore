package db

import (
	"database/sql"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
)

// Customer struct represents the structure of a Customer record in the database.
type Customer struct {
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
func (r *CustomerRepository) CreateCustomer(Customer *Customer) error {
	return nil
}

// GetCustomerByISBN retrieves a Customer from the database by its ISBN.
func (r *CustomerRepository) GetCustomerByID(id string) (*Customer, error) {
	return &Customer{}, nil
}

// UpdateCustomer updates an existing Customer record in the database.
func (r *CustomerRepository) UpdateCustomer(Customer *Customer) error {
	return nil
}

// DeleteCustomer removes a Customer from the database by its ISBN.
func (r *CustomerRepository) DeleteCustomer(id string) error {
	return nil
}

// GetAllCustomers retrieves all Customers from the database.
func (r *CustomerRepository) GetAllCustomers() ([]Customer, error) {
	return []Customer{}, nil
}
