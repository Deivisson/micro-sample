package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Driver to Postgresql connection
)

// Server struct to provide DB connection and Router
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize start database`s connection and routes
func (server *Server) Initialize() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// Run start the server
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 6000")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
