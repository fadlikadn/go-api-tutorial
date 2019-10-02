package middlewares

import (
	"errors"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"net/http"
	"github.com/fadlikadn/go-api-tutorial/api/auth"
)

/**
SetMiddlewareJSON: This will format all responses to JSON
SetMiddlewareAuthentication: This will check for the validity of the authentication token provided.
 */

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w,r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

