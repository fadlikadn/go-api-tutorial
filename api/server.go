package api

import (
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api/controllers"
	"github.com/fadlikadn/go-api-tutorial/api/seed"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	server = controllers.Server{}
)

func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not coming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(os.Getenv("DB_DRIVER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))

	APP_INSTALL, err := strconv.ParseBool(os.Getenv("APP_INSTALL_NEW"))
	if err != nil {
		log.Fatalf("Error getting env value APP_INSTALL_NEW")
	}

	if APP_INSTALL == true {
		seed.Load(server.DB)
	}

	//seed.MigrateOnly(server.DB)

	server.Run(os.Getenv("APP_PORT"))
}
