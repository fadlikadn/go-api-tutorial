package main

import (
	"flag"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/api"
	"github.com/fadlikadn/go-api-tutorial/api/controllers"
	"github.com/fadlikadn/go-api-tutorial/configuration"
	"github.com/fadlikadn/go-api-tutorial/persistence/dblayer"
	"log"
)

func main() {
	api.Run()

	/**
		Using new architecture based on book Microservice architecture
	 */
	confPath := flag.String("conf", `.\configuration\config.json`, "flag to set the path to the configuration json file" )
	flag.Parse()

	// extract configuration
	config, _ := configuration.ExtractConfiguration(*confPath)

	fmt.Println("Connecting to database")
	dbhandler, _ := dblayer.NewPersistenceLayer(config.Databasetype, config.DBConnection)
	// RESTful event API start
	log.Fatal(controllers.ServeAPI(config.RestfulEndpoint, dbhandler))
}
