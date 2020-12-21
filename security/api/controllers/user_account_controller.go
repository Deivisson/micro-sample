package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/deivisson/micro-sample/security/api/models"
	"github.com/deivisson/micro-sample/security/api/responses"
	"github.com/deivisson/micro-sample/security/api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) signUp(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	bodyParams, err := getBodyParams(w, r.Body)
	if err != nil {
		return
	}
	user.Create(server.DB, bodyParams)
	responses.Dispatch(w, user)
}

func (server *Server) signIn(w http.ResponseWriter, r *http.Request) {
	bodyParams, err := getBodyParams(w, r.Body)
	if err != nil {
		return
	}

	user := models.User{}
	if err = json.Unmarshal(bodyParams, &user); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err := user.ValidateSignIn(); err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := server.verifyUser(user.Email, user.Password)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.Success(w, http.StatusOK, token)
}

func (server *Server) verifyUser(email, password string) (interface{}, error) {
	var err error
	type result struct {
		Token string `json:"token"`
		Name  string `json:"name"`
	}
	user := models.User{}

	err = server.DB.Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", errors.New("Usuário ou Senha Inválidos")
	}

	err = verifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", errors.New("Usuário ou Senha Inválidos")
	}
	token, _ := utils.CreateToken(user.ID)
	return result{Token: token, Name: user.Name}, nil
}

func verifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (server *Server) emailAvailable(w http.ResponseWriter, r *http.Request) {
	var err error
	user := models.User{}
	log.Println(">>>>>>>>>> emailAvailable")
	email, err := getBodyParamByName(r, "email")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
	} else if email == nil || len(strings.TrimSpace(email.(string))) == 0 {
		errors := models.UserErrors{Email: []string{"Email deve ser informado"}}
		responses.Dispatch(w, errors)
	} else {
		finded, _ := user.FindByEmail(server.DB, email.(string))
		responses.Dispatch(w, !finded)
	}
}
