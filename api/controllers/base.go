package controllers

import (
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql database driver
	"github.com/rs/cors"
	"log"
	"net/http"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (server *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	if Dbdriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
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

	server.DB.Debug().AutoMigrate(&models.User{}, &models.Post{}) // database migration

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

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
		AllowCredentials: true,
	})
	handler := c.Handler(server.Router)
	log.Fatal(http.ListenAndServe(addr, handler))

}
