package controllers

import (
	"net/http"

	"github.com/deivisson/micro-sample/auth/api/responses"
)

func (server *Server) check(w http.ResponseWriter, r *http.Request) {
	responses.Success(w, http.StatusUnauthorized, nil)
}
