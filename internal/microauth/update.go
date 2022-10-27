package microauth

import (
	"fmt"
	"net/http"
)

func update(username string) UserCallback {
	return func(user User) (int, []byte, error) { 
		return http.StatusInternalServerError, []byte{}, fmt.Errorf("not yet implemented") 
	}
}
