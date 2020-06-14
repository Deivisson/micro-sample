package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/deivisson/apstore/api/models"
	"github.com/deivisson/apstore/api/responses"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // Driver to Postgresql connection
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	}
	server.DB.Debug().AutoMigrate(&models.User{}) //database migration
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func GetHttpParamsAsInt64(w http.ResponseWriter, r *http.Request, name string) (uint64, error) {
	value, err := strconv.ParseUint(mux.Vars(r)[name], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return 0, err
	}
	return uint64(value), nil
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
