package controllers

import (
	"github.com/fadlikadn/go-api-tutorial/api/middlewares"
	"net/http"
)

const (
	ASSETS_DIR = "/assets/"
	NODE_MODULES_DIR = "/node_modules/"
)

func (s *Server) initializeRoutes() {
	/**
	Web Routes
	 */
	s.Router.HandleFunc("/", s.HomeWeb).Methods("GET")
	s.Router.HandleFunc("/login", s.LoginWeb).Methods("GET")
	s.Router.HandleFunc("/register", s.RegisterWeb).Methods("GET")
	s.Router.HandleFunc("/forgotpassword", s.ForgotPasswordWeb).Methods("GET")
	s.Router.HandleFunc("/activation-pending", s.ActivationPending).Methods("GET")
	s.Router.HandleFunc("/users", s.ManageUserWeb).Methods("GET")

	/**
	Static Files such as JS, CSS, others
	 */
	s.Router.
		PathPrefix(ASSETS_DIR).
		Handler(http.StripPrefix(ASSETS_DIR, http.FileServer(http.Dir("."+ASSETS_DIR))))
	s.Router.
		PathPrefix(NODE_MODULES_DIR).
		Handler(http.StripPrefix(NODE_MODULES_DIR, http.FileServer(http.Dir("."+NODE_MODULES_DIR))))

	/**
	API Routes
	 */
	// Home Route
	s.Router.HandleFunc("/api", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/api/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")
	s.Router.HandleFunc("/api/register", middlewares.SetMiddlewareJSON(s.Register)).Methods("POST")
	s.Router.HandleFunc("/api/logout", middlewares.SetMiddlewareJSON(s.Logout)).Methods("GET")


	// Users routes
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/api/session/users", middlewares.SetMiddlewareJSON(s.CreateUserSession)).Methods("POST")
	s.Router.HandleFunc("/api/session/users", middlewares.SetMiddlewareJSON(s.GetUsersSession)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/api/session/users/{id}", middlewares.SetMiddlewareJSON(s.UpdateUserSession)).Methods("PUT")

	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")
	s.Router.HandleFunc("/api/session/users/{id}", s.DeleteUserSession).Methods("DELETE")

	// Posts routes
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")
}
