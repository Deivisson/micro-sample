package controllers

import (
	"github.com/deivisson/micro-sample/auth/api/middlewares"
)

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/verify", middlewares.SetMiddlewareJSON(s.check)).Methods("GET")
}
