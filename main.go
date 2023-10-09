package main

import (
	"log"
	"net/http"

	"github.com/siaudvytisViktoras/golang-sqlite-e-commerce/config"
	"github.com/siaudvytisViktoras/golang-sqlite-e-commerce/routes"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":"+config.AppConfig.ServerPort, nil))
}
