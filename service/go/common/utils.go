package common

import (
	"database/sql"
	"fmt"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

// InitDB initializes and returns a MySQL database connection.
func InitDB(config DBConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
