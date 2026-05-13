package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userId uuid.UUID, role string) (string, error) {
	tokenSecret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   userId,
		"role": role,
		"iat":  time.Now(),
		"eat":  time.Now().Add(time.Hour),
	})
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (uuid.UUID, string, error) {
	tokenSecret := os.Getenv("JWT_SECRET")
	myClaims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &myClaims, func(token *jwt.Token) (any, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, "", err
	}
	expiresAt, err := time.Parse(time.RFC3339, myClaims["eat"].(string))
	if err != nil {
		return uuid.Nil, "", err
	}
	if expiresAt.Before(time.Now()) {
		return uuid.Nil, "", errors.New("token expired")
	}
	userId, err := uuid.Parse(myClaims["id"].(string))
	if err != nil {
		return uuid.Nil, "", err
	}
	return userId, myClaims["role"].(string), nil
}

func MakeRefreshToken() string {
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return ""
	}
	token := hex.EncodeToString(tokenBytes)
	return token
}
