package microauth

import (
	"fmt"
	"net/http"
)

func delete(username string) (int, error) {
	return http.StatusInternalServerError, fmt.Errorf("not yet implemented")
}
