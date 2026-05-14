package main

import (
	"database/sql"
	"errors"
	"log"
	"os"

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
		log.Fatalf("Error loading .env file: %v\n", err)
		return
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
		return
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v\n", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing database connection: %v\n", err)
		}
	}()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database connection: %v\n", err)
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

	serveMux.HandleFunc("DELETE /api/products/{id}", apiCfg.DeleteProduct)

	serveMux.HandleFunc("POST /auth/register", apiCfg.Register)

	serveMux.HandleFunc("POST /auth/login", apiCfg.Login)

	serveMux.HandleFunc("POST /auth/logout", apiCfg.Logout)

	serveMux.HandleFunc("POST /auth/refresh", apiCfg.Refresh)

	server := &http.Server{
		Addr:    ":8080",
		Handler: corsMiddleware(serveMux),
	}

	log.Println("Server starting on port 8080")

	if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Error starting server: %v\n", err)
	} else {
		log.Println("Server stopped")
	}
}
