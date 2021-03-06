package api

import (
	"fmt"
	"log"
	"os"

	"github.com/deivisson/micro-sample/security/api/controllers"
	"github.com/deivisson/micro-sample/security/api/db/migrate"
	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Run is the start point of project
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	migrate.Load(server.DB)
	server.Run(":6000")
}
