package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
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
// func Error(w http.ResponseWriter, statusCode int, err interface{}) {
// 	if err != nil {
// 		Success(w, statusCode, getMessageError(err))
// 		return
// 	}
// 	Success(w, http.StatusBadRequest, nil)
// }
