package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	author_service "github.com/mayureshucsb2019/bookstore/service/author/service"
	book_service "github.com/mayureshucsb2019/bookstore/service/book/service"
	"github.com/mayureshucsb2019/bookstore/service/common"
	customer_service "github.com/mayureshucsb2019/bookstore/service/customer/service"
	"github.com/mayureshucsb2019/bookstore/service/factory"
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
	configFile := flag.String("config", "config.json", "path to the configuration file")
	config, err := loadConfig(*configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize DB connection
	dbConn, err := common.GetDBInstance(common.DBConfig{
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

	// Get the repository factory
	repoFactory := factory.GetRepositoryFactory(dbConn)

	// Create the book repository with the DB connection
	bookRepo := repoFactory.CreateBookRepository()
	bookAPIService := book_service.NewDefaultAPIService(bookRepo)
	bookAPIController := book_service.NewDefaultAPIController(bookAPIService)

	// Create the author repository with the DB connection
	authorRepo := repoFactory.CreateAuthorRepository()
	authorAPIService := author_service.NewDefaultAPIService(authorRepo)
	authorAPIController := author_service.NewDefaultAPIController(authorAPIService)

	// Create the author repository with the DB connection
	customerRepo := repoFactory.CreateCustomerRepository()
	customerAPIService := customer_service.NewDefaultAPIService(customerRepo)
	customerAPIController := customer_service.NewDefaultAPIController(customerAPIService)

	log.Printf("Server started")
	router := common.NewRouter(bookAPIController, authorAPIController, customerAPIController)

	log.Fatal(http.ListenAndServe(":8080", router))
}
