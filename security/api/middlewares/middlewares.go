package middlewares

import (
	"errors"
	"net/http"

	"github.com/deivisson/micro-sample/security/api/responses"
	"github.com/deivisson/micro-sample/security/api/utils"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetSecurityRoutes(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		err := utils.TokenValid(r)
		if err != nil {
			responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
