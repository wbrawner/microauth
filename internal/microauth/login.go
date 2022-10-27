package microauth

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func login(user User) (int, []byte, error) {
	passwordBytes := []byte(user.Password)
	for _, u := range users {
		if bcrypt.CompareHashAndPassword([]byte(u.Password), passwordBytes) == nil {
			return http.StatusAccepted, []byte{}, nil
		}
	}
	// TODO: Generate a session token, persist in database, and return in response
	return http.StatusUnauthorized, []byte{}, fmt.Errorf("credentials invalid")
}
