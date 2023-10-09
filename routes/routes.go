package routes

import (
	"github.com/siaudvytisViktoras/golang-sqlite-e-commerce/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("test", handlers.MyTestHandler).Methods("GET")
	// Define other routes and handlers
}
