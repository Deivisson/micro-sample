package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

	"github.com/deivisson/micro-sample/security/api/models"
)

// Success response to client
func Success(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(struct {
		Content interface{} `json:"content"`
	}{data})
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// Error process the failure request for especifics http codes, like 4xx, 5xx
func Error(w http.ResponseWriter, statusCode int, err interface{}) {
	if err != nil {
		Success(w, statusCode, getMessageError(err))
		return
	}
	Success(w, http.StatusBadRequest, nil)
}

// Dispatch response the httpServer according to data
func Dispatch(w http.ResponseWriter, data interface{}) {
	errors := getErrorFromModels(data)
	if errors.Exception != nil {
		Error(w, http.StatusInternalServerError, errors.Exception)
	} else if errors.RecordNotFound != nil {
		Error(w, http.StatusNotFound, errors.RecordNotFound)
	} else if errors.Business != nil {
		Error(w, http.StatusBadRequest, errors.Business)
	} else {
		Success(w, http.StatusOK, data)
	}
}

func getMessageError(errorsInterface interface{}) interface{} {
	switch reflect.TypeOf(errorsInterface).Kind() {
	case reflect.Struct:
		return struct {
			Errors interface{} `json:"errors"`
		}{errorsInterface}
	default:
		return struct {
			Message string `json:"message"`
		}{errorsInterface.(error).Error()}
	}
}

func getErrorFromModels(data interface{}) models.Errors {
	switch data.(type) {
	case models.User:
		return data.(models.User).Errors
	default:
		return models.Errors{}
	}
}
