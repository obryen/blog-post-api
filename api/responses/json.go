package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ToJsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

func toError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		toJsonResponse(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	toJsonResponse(w, http.StatusBadRequest, nil)
}
