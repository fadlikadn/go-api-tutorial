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
	"sync"
)

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	/*session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/

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

	// Create Session
	//session, _ := store.Get(r, "cookie-name")
	//session.Values["authenticated"] = true
	//_ = session.Save(r, w)

	// Using Gorilla Session
	/*session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}*/
	// Using SCS Session Manager
	//sessionManager.Put(r.Context(), "authenticated", true)

	// Using SecureCookie
	server.setSession(user.Email, w)

	//http.Redirect(w, r, "/", http.StatusFound)

	//var wg sync.WaitGroup
	//for i := 0; i < 1; i++ {
	//	wg.Add(1)
	//	go server.CreateSession(&wg, w, r)
	//}
	//
	//wg.Wait()

	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) CreateSession(wg *sync.WaitGroup, w http.ResponseWriter, r *http.Request) {
	defer wg.Done()

	session, _ := store.Get(r, "cookie-name")
	session.Values["authenticated"] = true
	_ = session.Save(r, w)
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
