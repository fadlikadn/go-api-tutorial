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
	username := server.getUsername(r)
	if username != "" {

	} else {
		http.Redirect(w, r,"/login", 302)
	}

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

func (server *Server) ActivationPending(w http.ResponseWriter, r *http.Request) {
	username := server.getUsername(r)
	if username != "" {
		http.Redirect(w, r, "/", 302)
	}

	var filepath = path.Join("views", "activation-pending.html")
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

func (server *Server) LoginWeb(w http.ResponseWriter, r *http.Request) {
	username := server.getUsername(r)
	if username != "" {
		http.Redirect(w, r, "/", 302)
	}

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
	username := server.getUsername(r)
	if username != "" {
		http.Redirect(w, r, "/", 302)
	}

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
	username := server.getUsername(r)
	if username != "" {
		http.Redirect(w, r, "/", 302)
	}

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
