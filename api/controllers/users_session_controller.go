package controllers

import (
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/fadlikadn/go-api-tutorial/api/responses"
	"net/http"
)

func (server *Server) GetUsersSession(w http.ResponseWriter, r *http.Request) {
	username := server.getUsername(r)
	if username != "" {

	} else {
		http.Redirect(w, r,"/login", 302)
	}
	user := models.User{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}