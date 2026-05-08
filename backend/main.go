package main

import (
	"database/sql"
	"errors"
	"os"

	"fmt"
	"net/http"

	"github.com/CowRules/MinMaxMacros/backend/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	dbQueries *database.Queries
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		fmt.Println("DB_URL environment variable not set")
		return
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			fmt.Printf("Error closing database connection: %v\n", err)
		}
	}()
	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database connection: %v\n", err)
		return
	}
	dbQueries := database.New(db)
	apiCfg := apiConfig{
		dbQueries: dbQueries,
	}
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("GET /api/products", apiCfg.GetProducts)

	serveMux.HandleFunc("GET /api/products/categories", apiCfg.GetProductCategories)

	serveMux.HandleFunc("GET /api/products/shops", apiCfg.GetProductShops)

	serveMux.HandleFunc("POST /api/products", apiCfg.CreateProduct)

	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(serveMux),
	}

	fmt.Println("Server starting on port 8080")

	if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Error starting server: %v\n", err)
	} else {
		fmt.Println("Server stopped")
	}
}
