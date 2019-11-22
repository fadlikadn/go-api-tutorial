package configuration

import (
	"encoding/json"
	"fmt"
	"github.com/fadlikadn/go-api-tutorial/persistence/dblayer"
	"os"
)

var (
	DBTypeDefault	=	dblayer.DBTYPE("mongodb")
	DBConnectionDefault = "mongodb://127.0.0.1"
	RestfulEPDefault = "localhost:8080"
)

type ServiceConfig struct {
	Databasetype dblayer.DBTYPE `json:"databasetype"`
	DBConnection string `json:"db_connection"`
	RestfulEndpoint string `json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		Databasetype:	DBTypeDefault,
		DBConnection:	DBConnectionDefault,
		RestfulEndpoint: RestfulEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	return conf, err
}
