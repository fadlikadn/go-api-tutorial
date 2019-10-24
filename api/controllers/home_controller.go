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
	//var filepath = path.Join("views", "index.html")
	//var tmpl, err = template.ParseFiles(filepath)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	indexTemplate := append(mainTemplateString, path.Join("views", "index-template.html"))
	var tmpl = template.Must(template.ParseFiles(indexTemplate...))

	//var tmpl = template.Must(template.ParseFiles(
	//	path.Join("views", "index-template.html"),
	//	path.Join("views", "_header.html"),
	//	path.Join("views", "_top-navbar.html"),
	//	path.Join("views", "_sidebar.html"),
	//	path.Join("views", "_content.html"),
	//	path.Join("views", "_footer.html"),
	//	path.Join("views", "_modals.html"),
	//	path.Join("views", "_js.html"),
	//))

	var data = M{
		"title": "Learning Golang",
		"name": "Mitrais",
		"sidebar": "home",
	}

	//err := tmpl.Execute(w, data)
	err := tmpl.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) ActivationPending(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "activation-pending.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data = M{}

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
