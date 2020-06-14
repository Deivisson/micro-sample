package controllers

import "github.com/deivisson/apstore/api/middlewares"

func (s *Server) initializeRoutes() {

	// Account Routes
	s.Router.HandleFunc("/users/account/sign-up", middlewares.SetMiddlewareJSON(s.signUp)).Methods("POST")
	s.Router.HandleFunc("/users/account/sign-in", middlewares.SetMiddlewareJSON(s.signIn)).Methods("POST")
	s.Router.HandleFunc("/users/account/email-available", middlewares.SetMiddlewareJSON(s.emailAvailable)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users/search", middlewares.SetSecurityRoutes(s.getAll)).Methods("POST")
	s.Router.HandleFunc("/users/{id}", middlewares.SetSecurityRoutes(s.getUserID)).Methods("GET")
	s.Router.HandleFunc("/users/avatar", middlewares.SetSecurityRoutes(s.uploadUserAvatar)).Methods("POST")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetSecurityRoutes(s.updateUser)).Methods("PUT")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetSecurityRoutes(s.deleteUser)).Methods("DELETE")
	// s.Router.HandleFunc("/users/{id}", middlewares.SetSecurityRoutes(s.deleteUser)).Methods("DELETE")

	//Stores routes
	// s.Router.HandleFunc("/stores/search", middlewares.SetSecurityRoutes(s.searchStores)).Methods("POST")
	// s.Router.HandleFunc("/stores/{id}", middlewares.SetSecurityRoutes(s.getStoreByID)).Methods("GET")
	// s.Router.HandleFunc("/stores", middlewares.SetSecurityRoutes(s.createStore)).Methods("POST")
	// s.Router.HandleFunc("/stores/name-available", middlewares.SetMiddlewareJSON(s.nameAvailable)).Methods("POST")
}
