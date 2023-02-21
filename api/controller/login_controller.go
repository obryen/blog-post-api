package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/obryen/blog-api/api/auth"
	"github.com/obryen/blog-api/api/models"
	"github.com/obryen/blog-api/api/responses"
	formaterror "github.com/obryen/blog-api/api/utils"
	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ToError(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.Formarterror(err.Error())
		responses.ToError(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.ToJsonResponse(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.GenerateToken(user.ID)
}
