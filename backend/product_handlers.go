package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/CowRules/MinMaxMacros/backend/internal/auth"
	"github.com/CowRules/MinMaxMacros/backend/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := cfg.dbQueries.GetProducts(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(products)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		log.Println(err)
		return
	}
}

func (cfg *apiConfig) GetProductCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := cfg.dbQueries.GetProductCategories(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	data, err := json.Marshal(categories)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		log.Println(err)
		return
	}
}

func (cfg *apiConfig) CreateProduct(w http.ResponseWriter, r *http.Request) {
	authenticated, userId, _ := auth.IsAuthenticated(r)
	if !authenticated {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	requestBody := database.CreateProductParams{}
	if err := decoder.Decode(&requestBody); err != nil {
		log.Println(err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	requestBody.UserID = userId
	if err := ValidateCreateProduct(requestBody); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := cfg.dbQueries.CreateProduct(r.Context(), requestBody)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(product)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(data); err != nil {
		log.Println(err)
		return
	}
}

func (cfg *apiConfig) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	authenticated, userId, _ := auth.IsAuthenticated(r)
	if !authenticated {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	productId, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		log.Println(err)
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	product, err := cfg.dbQueries.GetProduct(r.Context(), productId)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "product does not exist", http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if userId != product.UserID {
		http.Error(w, "unauthorized", http.StatusForbidden)
		return
	}
	if err := cfg.dbQueries.DeleteProduct(r.Context(), productId); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (cfg *apiConfig) GetProductShops(w http.ResponseWriter, r *http.Request) {
	shops, err := cfg.dbQueries.GetProductShops(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(shops)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		log.Println(err)
		return
	}
}
