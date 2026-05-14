package auth

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func IsAuthenticated(r *http.Request) (bool, uuid.UUID, string) {
	cookie, err := r.Cookie("AccessToken")
	if errors.Is(err, http.ErrNoCookie) {
		log.Println("No AccessToken Cookie")
		return false, uuid.Nil, ""
	}
	userId, role, err := ValidateToken(cookie.Value)
	if err != nil {
		log.Println(err)
		return false, uuid.Nil, ""

	}
	return true, userId, role
}
