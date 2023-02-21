package controller

import (
	"net/http"

	"github.com/obryen/blog-api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.ToJsonResponse(w, http.StatusOK, "Welcome To This Awesome API")
}
