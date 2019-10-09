package controllers

import (
	"encoding/json"
	"errors"
	"github.com/fadlikadn/go-api-tutorial/api/auth"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//responses.ERROR(w, http.StatusUnprocessableEntity, err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error when read r.Body"))
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		//responses.ERROR(w, http.StatusUnprocessableEntity, err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error when process json.Unmarshal"))
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		//responses.ERROR(w, http.StatusUnprocessableEntity, err)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error when validate users"))
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		//formattedError := formateerror.FormatError(err.Error())
		//responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error when format error"))
		return
	}
	responses.JSON(w, http.StatusOK, token)
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
	return auth.CreateToken(user.ID)
}
