package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/deivisson/micro-sample/security/api/responses"
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
func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		DbHost, DbPort, DbUser, DbName, DbPassword,
	)
	if server.DB, err = gorm.Open(Dbdriver, connectionString); err != nil {
		if err = createDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName); err != nil {
			fmt.Printf("Cannot create database %s", DbName)
			log.Fatal("This is the error:", err)
		}
		server.DB, err = gorm.Open(Dbdriver, connectionString)
	}
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// Run start the server
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 6000")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

// GetHTTPParamsAsInt64 return int params
func GetHTTPParamsAsInt64(w http.ResponseWriter, r *http.Request, name string) (uint64, error) {
	value, err := strconv.ParseUint(mux.Vars(r)[name], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return 0, err
	}
	return uint64(value), nil
}

func createDatabase(Dbdriver, DbUser, DbPassword, DbPort, DbHost, dbName string) error {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable dbname=postgres", DbHost, DbPort, DbUser, DbPassword)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		return err
	}

	if db := db.Exec(fmt.Sprintf("CREATE DATABASE %s;", dbName)); db.Error != nil {
		fmt.Println("Error on create database")
		return err
	}
	return nil
}

func getBodyParamByName(r *http.Request, paramName string) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var t interface{}

	if err := decoder.Decode(&t); err != nil {
		return nil, err
	}
	mapParms := t.(map[string]interface{})
	return mapParms[paramName], nil
}

func getBodyParams(w http.ResponseWriter, reader io.Reader) ([]byte, error) {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return nil, err
	}
	return body, nil
}
