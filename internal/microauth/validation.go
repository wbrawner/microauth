package microauth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func validateApiKey(apiKey string, callback http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("X-API-Key")
		if apiKey != authHeader {
			http.Error(w, "invalid api key", http.StatusForbidden)
		} else {
			callback(w, r)
		}
	}
}

func validateHttpMethod(allowedMethod string, callback http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != allowedMethod {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		} else {
			callback(w, r)
		}
	}
}

func validateUserBody(callback UserCallback) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			log.Printf("error: failed to read request body: %v", err)
			return
		}

		var user User
		err = json.Unmarshal(body, &user)
		if err != nil {
			http.Error(w, "invalid request body", http.StatusBadRequest)
			log.Printf("error: failed to unmarshal login request: %v", err)
			return
		}

		returnCode, responseBody, err := callback(user)
		
		if err != nil {
			http.Error(w, fmt.Sprintf("%v", err), returnCode)
			return
		}

		w.WriteHeader(returnCode)
		w.Write(responseBody)
	}
}
