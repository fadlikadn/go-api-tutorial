package controllers

import (
	"github.com/fadlikadn/go-api-tutorial/api/middlewares"
	"github.com/fadlikadn/go-api-tutorial/persistence"
	"github.com/gorilla/mux"
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
	// Auth
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareAuthenticationSessionOut(s.LoginWeb)).Methods("GET")
	s.Router.HandleFunc("/register", middlewares.SetMiddlewareAuthenticationSessionOut(s.RegisterWeb)).Methods("GET")
	s.Router.HandleFunc("/forgotpassword", middlewares.SetMiddlewareAuthenticationSessionOut(s.ForgotPasswordWeb)).Methods("GET")
	s.Router.HandleFunc("/activation-pending", s.ActivationPending).Methods("GET")

	// Protobuf + GRPC
	s.Router.HandleFunc("/grpc/test-protobuf", middlewares.SetMiddlewareAuthenticationSessionOut(s.TestProtobuf)).Methods("GET")

	// Features
	s.Router.HandleFunc("/", s.IndexMain).Methods("GET")
	s.Router.HandleFunc("/dashboard", middlewares.SetMiddlewareAuthenticationSession(s.DashboardHome)).Methods("GET")
	s.Router.HandleFunc("/dashboard/users", middlewares.SetMiddlewareAuthenticationSession(s.ManageUserWeb)).Methods("GET")
	s.Router.HandleFunc("/dashboard/customers", middlewares.SetMiddlewareAuthenticationSession(s.ManageCustomerWeb)).Methods("GET")
	s.Router.HandleFunc("/dashboard/service-transactions", middlewares.SetMiddlewareAuthenticationSession(s.ManageServiceTransactionWeb)).Methods("GET")
	s.Router.HandleFunc("/dashboard/add-service-transactions", middlewares.SetMiddlewareAuthenticationSession(s.AddServiceTransactionWeb)).Methods("GET")
	s.Router.HandleFunc("/dashboard/map-basic", middlewares.SetMiddlewareAuthenticationSession(s.LeafletBasicMap)).Methods("GET")

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

	// Customers routes
	s.Router.HandleFunc("/api/customers", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.GetCustomers))).Methods("GET")
	s.Router.HandleFunc("/api/customers", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.CreateCustomer))).Methods("POST")
	s.Router.HandleFunc("/api/customers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.GetCustomer))).Methods("GET")
	s.Router.HandleFunc("/api/customers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.UpdateCustomer))).Methods("PUT")
	s.Router.HandleFunc("/api/customers/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.DeleteCustomer))).Methods("DELETE")

	// Service Transaction routes
	s.Router.HandleFunc("/api/service-transactions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.GetServiceTransactions))).Methods("GET")
	s.Router.HandleFunc("/api/service-transactions", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.CreateServiceTransaction))).Methods("POST")
	s.Router.HandleFunc("/api/service-transactions-complex", middlewares.SetMiddlewareJSON(s.CreateComplexServiceTransaction)).Methods("POST")
	s.Router.HandleFunc("/api/service-transactions-complex/{id}", middlewares.SetMiddlewareJSON(s.UpdateComplexServiceTransaction)).Methods("PUT")
	s.Router.HandleFunc("/api/service-transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.GetServiceTransaction))).Methods("GET")
	s.Router.HandleFunc("/api/service-transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.UpdateServiceTransaction))).Methods("PUT")
	s.Router.HandleFunc("/api/service-transactions/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.DeleteServiceTransaction))).Methods("DELETE")
	s.Router.HandleFunc("/api/service-transactions/sendstatusemail/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.SendTransactionStatusEmail))).Methods("GET")
	s.Router.HandleFunc("/api/invoice/sample-service-transactions", middlewares.SetMiddlewareAuthenticationSession(s.CreateSampleInvoiceServiceTransaction)).Methods("GET")
	s.Router.HandleFunc("/api/invoice/service-transactions/{uuid}", middlewares.SetMiddlewareAuthenticationSession(s.CreateInvoiceServiceTransaction)).Methods("GET")
	s.Router.HandleFunc("/api/dashboard/service-transaction", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthenticationSession(s.GetServicesAmount))).Methods("GET")
	// Posts routes
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.CreatePost)).Methods("POST")
	s.Router.HandleFunc("/api/posts", middlewares.SetMiddlewareJSON(s.GetPosts)).Methods("GET")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(s.GetPost)).Methods("GET")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdatePost))).Methods("PUT")
	s.Router.HandleFunc("/api/posts/{id}", middlewares.SetMiddlewareAuthentication(s.DeletePost)).Methods("DELETE")

	// Search Invoice
	s.Router.HandleFunc("/api/search/invoice/{invoiceno}", middlewares.SetMiddlewareJSON(s.SearchInvoice)).Methods("GET")
}

// TODO Implement Technician Module
// TODO Implement Reporting Status for each Service Transaction
// TODO User's Level

// TODO Implement SAAS Starter

//
/**
	Microservice Event API
	Using interface and abstraction
 */
func ServeAPI(endpoint string, databasehandler persistence.DatabaseHandler) error {
	handler := NewEventHandler(databasehandler)
	r := mux.NewRouter()
	eventsRouter := r.PathPrefix("/events").Subrouter()
	eventsRouter.Methods("GET").Path("/{SearchCriteria}/{search}").HandlerFunc(handler.FindEventHandler)
	eventsRouter.Methods("GET").Path("").HandlerFunc(handler.AllEventHandler)
	eventsRouter.Methods("POST").Path("").HandlerFunc(handler.NewEventHandler)

	return http.ListenAndServe(endpoint, r)
}
