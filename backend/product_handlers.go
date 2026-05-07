package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *apiConfig) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := cfg.dbQueries.GetProducts(r.Context())
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(products)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
	data, err := json.Marshal(categories)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	data, err := json.Marshal(shops)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(data); err != nil {
		fmt.Println(err)
		return
	}
}
