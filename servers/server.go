package server

import (
	"fmt"
	"log"
	"net/http"
	"github.com/BerylCAtieno/basic-information-api/routes"
)

func StartServer() {
	router := routes.RegisterRoutes()
	port := "8080"

	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
