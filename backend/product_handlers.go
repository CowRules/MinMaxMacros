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
