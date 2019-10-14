package controllers

import (
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"html/template"
	"net/http"
	"path"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this awesome API")
}

func (server *Server) HomeWeb(w http.ResponseWriter, r *http.Request) {
	/*session, _ := store.Get(r, "cookie-name")

	if session.Values["authenticated"] != true {
		http.Redirect(w, r, base_url + "/login", 301)
	}*/
	session, ok := sessionManager.Get(r.Context(), "authenticated").(bool)
	if !session || !ok {
		http.Redirect(w, r, base_url + "/login", 301)
	}

	// Check if usr is authenticated
	/*if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		// redirect to login
		http.Redirect(w, r, base_url + "/login", 301)
	}*/

	var filepath = path.Join("views", "index.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{} {
		"title": "Learning Golang",
		"name": "Mitrais",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) LoginWeb(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "login.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{} {}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) RegisterWeb(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "register.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{} {}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) ForgotPasswordWeb(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "forgot-password.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = map[string]interface{} {}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
