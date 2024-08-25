package common

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

type DBConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
}

type DBConnection struct {
	*sql.DB
}

var dbInstance *DBConnection
var once sync.Once

func GetDBInstance(config DBConfig) (*DBConnection, error) {
	once.Do(func() {
		connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
		db, err := sql.Open("mysql", connStr)
		if err != nil {
			log.Fatalf("Failed to connect to the database: %v", err)
		}

		dbInstance = &DBConnection{
			DB: db,
		}
	})

	// Test the connection
	if err := dbInstance.DB.Ping(); err != nil {
		log.Fatalf("Database connection error %+v", err)
		return nil, err
	}
	log.Printf("Database connection is successful")

	return dbInstance, nil
}
