package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	routes "github.com/mrandycodes/kyber/internal"
	"github.com/mrandycodes/kyber/internal/platform/storage/in_memory"
)

const httpAddr = ":8080"

var repository routes.RoutesRepository

type addRequest struct {
	Route string `json:"route"`
}

type deleteRouteRequest struct {
	Route string `json:"route"`
}

type routesResponse struct {
	Route string `json:"route"`
}

func main() {
	fmt.Println("Starting server on port", httpAddr)

	repository = in_memory.NewRoutesRepository()

	mux := mux.NewRouter()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/routes", addRouteHandler).Methods("POST")
	mux.HandleFunc("/routes", deleteRouteHandler).Methods("DELETE")
	mux.HandleFunc("/routes", listRoutesHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(httpAddr, mux))
}

func healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("everything is ok!"))
}

func addRouteHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var addReq addRequest
	json.NewDecoder(req.Body).Decode(&addReq)

	route, err := routes.NewRoute(addReq.Route)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	err = repository.Add(route)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusCreated)
}

func deleteRouteHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var deleteRouteReq deleteRouteRequest
	json.NewDecoder(req.Body).Decode(&deleteRouteReq)

	route, err := routes.NewRoute(deleteRouteReq.Route)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	err = repository.Delete(route)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	}

	res.WriteHeader(http.StatusAccepted)
}

func listRoutesHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	response := []routesResponse{}
	routes := repository.List()

	for _, route := range routes {
		response = append(response, routesResponse{route.Value()})
	}

	json.NewEncoder(res).Encode(response)
}
