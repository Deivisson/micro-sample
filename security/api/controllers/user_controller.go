package controllers

import (
	"net/http"

	"github.com/deivisson/apstore/api/models"
	"github.com/deivisson/apstore/api/responses"
)

func (server *Server) getAll(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	users, err := user.FindAll(server.DB)

	if err == nil {
		responses.Success(w, http.StatusOK, users)
	} else {
		responses.Error(w, http.StatusInternalServerError, err)
	}
}

func (server *Server) getUserID(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	id, err := GetHttpParamsAsInt64(w, r, "id")
	if err != nil {
		return
	}
	user.FindByID(server.DB, id)
	responses.Dispatch(w, user)
}

// func (server *Server) updateUser(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	uid, err := strconv.ParseUint(vars["id"], 10, 32)
// 	if err != nil {
// 		responses.Error(w, http.StatusBadRequest, err)
// 		return
// 	}
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		responses.Error(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	user := models.User{}
// 	err = json.Unmarshal(body, &user)
// 	if err != nil {
// 		responses.Error(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	tokenID, err := utils.ExtractTokenID(r)
// 	if err != nil {
// 		responses.Error(w, http.StatusUnauthorized, errors.New("Unauthorized"))
// 		return
// 	}
// 	if tokenID != uint32(uid) {
// 		responses.Error(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
// 		return
// 	}
// 	user.Prepare()
// 	err = user.Validate("update")
// 	if err != nil {
// 		responses.Error(w, http.StatusUnprocessableEntity, err)
// 		return
// 	}
// 	updatedUser, err := user.UpdateUser(server.DB, uint32(uid))
// 	if err != nil {
// 		formattedError := utils.FormatError(err.Error())
// 		responses.Error(w, http.StatusInternalServerError, formattedError)
// 		return
// 	}
// 	responses.Success(w, http.StatusOK, updatedUser)
// }

func (server *Server) deleteUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}

	id, err := GetHttpParamsAsInt64(w, r, "id")
	if err != nil {
		return
	}
	id, err = user.Delete(server.DB, id)
	responses.Dispatch(w, user)
}

func (server *Server) uploadUserAvatar(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	avatar, err := user.UploadAvatar(server.DB, r)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
	} else {
		responses.Dispatch(w, avatar)
	}
}
