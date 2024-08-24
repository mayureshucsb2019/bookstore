package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	author_db "github.com/mayureshucsb2019/bookstore/go/author/db"
	author_service "github.com/mayureshucsb2019/bookstore/go/author/service"
	book_db "github.com/mayureshucsb2019/bookstore/go/book/db"
	book_service "github.com/mayureshucsb2019/bookstore/go/book/service"
	"github.com/mayureshucsb2019/bookstore/go/common"
)

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"dbname"`
}

// LoadConfig reads the configuration from a JSON file.
func loadConfig(filePath string) (Config, error) {
	var config Config
	file, err := os.Open(filePath)
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func main() {

	// Load configuration from file
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize DB connection
	dbConn, err := common.InitDB(common.DBConfig{
		Username: config.Username,
		Password: config.Password,
		Host:     config.Host,
		Port:     config.Port,
		DBName:   config.DBName,
	})

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer dbConn.Close()

	// Test the connection
	if err := TestDBConnection(dbConn); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Printf("Database connection is successful")

	// Create the book repository with the DB connection
	bookRepo := book_db.NewBookRepository(dbConn)
	bookAPIService := book_service.NewDefaultAPIService(bookRepo)
	bookAPIController := book_service.NewDefaultAPIController(bookAPIService)

	// Create the author repository with the DB connection
	authorRepo := author_db.NewAuthorRepository(dbConn)
	authorAPIService := author_service.NewDefaultAPIService(authorRepo)
	authorAPIController := author_service.NewDefaultAPIController(authorAPIService)

	log.Printf("Server started")
	router := common.NewRouter(bookAPIController, authorAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// TestDBConnection tests the database connection
func TestDBConnection(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
