package main

import (
	"errors"

	"github.com/CowRules/MinMaxMacros/backend/internal/database"
)

func ValidateCreateProduct(body database.CreateProductParams) error {
	if body.Name == "" {
		return errors.New("missing product name")
	}
	if body.Calories < 0 || body.Protein < 0 || body.Fiber < 0 || body.Price < 0 || body.Weight < 0 {
		return errors.New("values cannot be negative")
	}
	return nil
}
