package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CowRules/MinMaxMacros/backend/internal/database"
)

func (cfg *apiConfig) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := cfg.dbQueries.GetProducts(r.Context())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		fmt.Println(err)
		return
	}
}

func (cfg *apiConfig) GetProductCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := cfg.dbQueries.GetProductCategories(r.Context())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	data, err := json.Marshal(categories)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		fmt.Println(err)
		return
	}
}

func (cfg *apiConfig) GetProductShops(w http.ResponseWriter, r *http.Request) {
	shops, err := cfg.dbQueries.GetProductShops(r.Context())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(shops)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		fmt.Println(err)
		return
	}
}

func (cfg *apiConfig) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	requestBody := database.CreateProductParams{}
	if err := decoder.Decode(&requestBody); err != nil {
		fmt.Println(err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := ValidateCreateProduct(requestBody); err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	product, err := cfg.dbQueries.CreateProduct(r.Context(), requestBody)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(data); err != nil {
		fmt.Println(err)
		return
	}
}
