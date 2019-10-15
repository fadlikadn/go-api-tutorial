package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/auth"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"github.com/fadlikadn/go-api-tutorial/api/utils/formateerror"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	user.IsActive = false
	err = user.Validate("")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	registeredUser, err := user.SaveUser(server.DB)

	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, registeredUser.ID))
	responses.JSON(w, http.StatusCreated, registeredUser)
}

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		//responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error when read r.Body"))
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		//responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error when process json.Unmarshal"))
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		//responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("error when validate users"))
		return
	}
	token, err := server.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formateerror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		//responses.ERROR(w, http.StatusUnprocessableEntity, errors.New("Error when format error"))
		return
	}

	// Using SecureCookie
	server.setSession(user.Email, w)

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("email = ? AND is_active = 1", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}

func (server *Server) Logout(w http.ResponseWriter, r *http.Request) {
	/*session, _ := store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	_ = session.Save(r, w)*/
	//sessionManager.Remove(r.Context(), "authenticated")
	server.clearSession(w);

	http.Redirect(w, r, base_url + "/login", 301)

	//responses.JSON(w, http.StatusOK, true)
}
