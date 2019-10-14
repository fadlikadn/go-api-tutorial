package main

import (
	"github.com/fadlikadn/go-api-tutorial/api"
	"github.com/fadlikadn/go-api-tutorial/api/controllers"
)

var (
	server = controllers.Server{}
)
func main() {
	api.Run()
}
