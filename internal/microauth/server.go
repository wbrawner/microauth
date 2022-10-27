package microauth

import (
	"fmt"
	"log"
	"net/http"
)

// TODO: Save this in a database, or maybe wrap in an interface to be able to use an in-memory store for tests
var users []User

func StartServer(port int, apiKey string) error {
	http.HandleFunc("/login", validateApiKey(apiKey, validateHttpMethod(http.MethodPost, validateUserBody(login))))
	http.HandleFunc("/register", validateApiKey(apiKey, validateHttpMethod(http.MethodPost, validateUserBody(register))))
	http.HandleFunc("/update", validateApiKey(apiKey, validateHttpMethod(http.MethodPut, func(w http.ResponseWriter, r *http.Request) {
		validateUserBody(update(r.URL.Query().Get("user")))(w, r)
	})))
	http.HandleFunc("/delete", validateApiKey(apiKey, validateHttpMethod(http.MethodDelete, func(w http.ResponseWriter, r *http.Request) {
		returnCode, err := delete(r.URL.Query().Get("user"))
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), returnCode)
			return
		}

		w.WriteHeader(returnCode)
		w.Write([]byte{})
	})))
	log.Printf("microauth server listening on :%d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
