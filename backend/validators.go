package main

import (
	"errors"
	"regexp"

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

func ValidateRegister(body database.CreateUserParams) error {
	isEmail, err := regexp.MatchString("^[^\\s]+@[a-zA-Z0-9-]+\\.[A-Za-z]+$", body.Email)
	if err != nil {
		return err
	}
	if !isEmail {
		return errors.New("invalid email")
	}
	if len(body.Password) < 8 || len(body.Password) > 100 {
		return errors.New("password must be between 8 and 100 characters")
	}
	if len(body.Username) < 5 || len(body.Username) > 20 {
		return errors.New("username must be between 5 and 20 characters")
	}
	isUsernameValid, err := regexp.MatchString("^[a-zA-Z0-9]+$", body.Username)
	if err != nil {
		return err
	}
	if !isUsernameValid {
		return errors.New("username can only contain alphanumeric characters")
	}
	return nil
}
