package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthMiddleware() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			http.Error(w, "Invalid Authorization Header", http.StatusUnauthorized)
			return
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			http.Error(w, "Invalid Credentials", http.StatusUnauthorized)
			return
		}

		username, password := credentials[0], credentials[1]
		if username == "admin" && password == "password" {
			// Proceed with the request
			return
		} else {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
	}
}
