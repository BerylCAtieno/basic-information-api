package routes

import (
	"github.com/gorilla/mux"
	"github.com/BerylCAtieno/basic-information-api/handlers"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetInfoHandler).Methods("GET")
	return router
}
