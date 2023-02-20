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

func ToError(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		ToJsonResponse(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	ToJsonResponse(w, http.StatusBadRequest, nil)
}
