package controllers

import (
	"github.com/deivisson/micro-sample/security/api/middlewares"
)

func (s *Server) initializeRoutes() {
	// Account Routes
	s.Router.HandleFunc("/users/account/sign-up", middlewares.SetMiddlewareJSON(s.signUp)).Methods("POST")
	s.Router.HandleFunc("/users/account/sign-in", middlewares.SetMiddlewareJSON(s.signIn)).Methods("POST")
	s.Router.HandleFunc("/users/account/email-available", middlewares.SetMiddlewareJSON(s.emailAvailable)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users/search", middlewares.SetSecurityRoutes(s.getAll)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetSecurityRoutes(s.getUserID)).Methods("GET")
}
