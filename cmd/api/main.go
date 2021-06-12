package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/mrandycodes/kyber/internal"
	"github.com/mrandycodes/kyber/internal/platform/server/handler/health"
	routes_handler "github.com/mrandycodes/kyber/internal/platform/server/handler/routes"
	"github.com/mrandycodes/kyber/internal/platform/storage/in_memory"
)

const httpAddr = ":8080"

var repository routes.RoutesRepository

type routesResponse struct {
	Route string `json:"route"`
}

func main() {
	fmt.Println("Starting server on port", httpAddr)

	repository = in_memory.NewRoutesRepository()

	mux := mux.NewRouter()
	mux.HandleFunc("/health", health.HealthHandler())
	mux.HandleFunc("/routes", routes_handler.AddRouteHandler(repository)).Methods("POST")
	mux.HandleFunc("/routes", routes_handler.DeleteRouteHandler(repository)).Methods("DELETE")
	mux.HandleFunc("/routes", routes_handler.ListRoutesHandler(repository)).Methods("GET")

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
