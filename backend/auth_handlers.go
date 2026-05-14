package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/CowRules/MinMaxMacros/backend/internal/auth"
	"github.com/CowRules/MinMaxMacros/backend/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) Register(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestBody := database.CreateUserParams{}
	if err := decoder.Decode(&requestBody); err != nil {
		log.Println(err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if err := ValidateRegister(requestBody); err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}
	_, err := cfg.dbQueries.GetUserByEmail(r.Context(), requestBody.Email)
	if err == nil {
		http.Error(w, "user with this email already exists", http.StatusConflict)
		return
	} else if !errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
	hashedPassword, err := auth.HashPassword(requestBody.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	requestBody.Password = hashedPassword
	if err := cfg.dbQueries.CreateUser(r.Context(), requestBody); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	return
}

func (cfg *apiConfig) Login(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	requestBody := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	if err := decoder.Decode(&requestBody); err != nil {
		log.Println(err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	user, err := cfg.dbQueries.GetUserByEmail(r.Context(), requestBody.Email)
	if errors.Is(err, sql.ErrNoRows) {
		log.Println(err)
		http.Error(w, "incorrect email or password", http.StatusUnauthorized)
		return
	}
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	isPasswordCorrect, err := auth.CheckPasswordHash(requestBody.Password, user.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	if !isPasswordCorrect {
		http.Error(w, "incorrect email or password", http.StatusUnauthorized)
		return
	}
	token, err := auth.MakeJWT(user.ID, user.Role)
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	refreshToken := auth.MakeRefreshToken()
	err = cfg.dbQueries.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{Token: refreshToken, UserID: user.ID, ExpiresAt: time.Now().Add(time.Hour * 24).UTC()})
	if err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    token,
		HttpOnly: true,
		Path:     "/api",
		SameSite: http.SameSiteLaxMode,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Path:     "/auth",
		SameSite: http.SameSiteLaxMode,
	})
	data, err := json.Marshal(struct {
		ID       uuid.UUID `json:"id"`
		Username string    `json:"username"`
		Role     string    `json:"role"`
	}{ID: user.ID, Username: user.Username, Role: user.Role})
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

func (cfg *apiConfig) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, err := r.Cookie("RefreshToken")
	if err != nil {
		log.Println(err)
		cfg.Logout(w, r)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	refreshToken, err := cfg.dbQueries.GetRefreshToken(r.Context(), refreshTokenCookie.Value)
	if err != nil {
		log.Println(err)
		cfg.Logout(w, r)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	if refreshToken.ExpiresAt.Before(time.Now()) || refreshToken.RevokedAt.Valid {
		cfg.Logout(w, r)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	user, err := cfg.dbQueries.GetUserById(r.Context(), refreshToken.UserID)
	if err != nil {
		log.Println(err)
		cfg.Logout(w, r)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	accessToken, err := auth.MakeJWT(user.ID, user.Role)
	if err != nil {
		log.Println(err)
		cfg.Logout(w, r)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    accessToken,
		HttpOnly: true,
		Path:     "/api",
		SameSite: http.SameSiteLaxMode,
	})
	w.WriteHeader(http.StatusOK)
}

func (cfg *apiConfig) Logout(w http.ResponseWriter, r *http.Request) {
	refreshTokenCookie, _ := r.Cookie("RefreshToken")
	if refreshTokenCookie != nil {
		_ = cfg.dbQueries.RevokeRefreshToken(r.Context(), refreshTokenCookie.Value)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    "",
		HttpOnly: true,
		Path:     "/auth",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    "",
		HttpOnly: true,
		Path:     "/api",
		SameSite: http.SameSiteLaxMode,
		MaxAge:   -1,
	})
}
