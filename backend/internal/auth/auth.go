package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func IsAuthenticated(r *http.Request) (bool, uuid.UUID, string) {
	cookie, err := r.Cookie("AccessToken")
	if errors.Is(err, http.ErrNoCookie) {
		fmt.Println("No AccessToken Cookie")
		return false, uuid.Nil, ""
	}
	userId, role, err := ValidateToken(cookie.Value)
	if err != nil {
		fmt.Println(err)
		return false, uuid.Nil, ""

	}
	return true, userId, role
}
