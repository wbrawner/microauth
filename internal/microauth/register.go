package microauth

import (
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func register(user User) (int, []byte, error) {
	for _, u := range users {
		if user.Name == u.Name {
			return http.StatusBadRequest, []byte{}, fmt.Errorf("username already taken")
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, []byte{}, fmt.Errorf("failed to hash password: %v", err)
	}
	users = append(users, User{Name: user.Name, Password: string(hashedPassword)})
	// TODO: Generate a session token, persist in database, and return in response
	return http.StatusNoContent, []byte{}, nil
}
