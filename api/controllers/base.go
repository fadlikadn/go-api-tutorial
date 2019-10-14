package controllers

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte(os.Getenv("API_SECRET"))
	store = sessions.NewCookieStore(key)
	base_url = "http://localhost:8080"
	sessionManager *scs.SessionManager
	cookieHandler = securecookie.New(
		securecookie.GenerateRandomKey(64),
		securecookie.GenerateRandomKey(32))
)

func init() {
	key = []byte(os.Getenv("API_SECRET"))
	store = sessions.NewCookieStore(key)
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	if Dbdriver == "mysql" {
		DBURL := connectionString
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	if Dbdriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DB, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", Dbdriver)
		}
	}

	/**
		Handle Session using SCS package
	 */
	// Initialize a new session manager
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	// sessionManager.Store = mysqlstore.New(server.DB.DB())

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}, &models.Session{}) // database migration

	server.Router = mux.NewRouter().StrictSlash(true)

	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	//log.Fatal(http.ListenAndServe(addr, server.Router))

	// Apply the CORS middleware to our top-level router, with the defaults.
	//log.Fatal(http.ListenAndServe(addr, handlers.CORS()(server.Router)))

	//corsObj := handlers.AllowedOrigins([]string{"*"})
	//log.Fatal(http.ListenAndServe(addr, handlers.CORS(corsObj)(server.Router)))

	// handle CORS using package cors
	/*c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(server.Router)
	log.Fatal(http.ListenAndServe(addr, handler))*/

	log.Fatal(http.ListenAndServe(addr, sessionManager.LoadAndSave(server.Router)))

}

func (server *Server) setSession(username string, w http.ResponseWriter) {
	value := map[string]string {
		"name": username,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name: "session",
			Value: encoded,
			Path: "/",
		}
		http.SetCookie(w, cookie)
	}
}

func (server *Server) getUsername(r *http.Request) (username string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			username = cookieValue["name"]
		}
	}
	return username
}

func (server *Server) clearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
}
