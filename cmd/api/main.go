package main

import (
	"encoding/json"
	"fmt"
	routes "github.com/mrandycodes/kyber/internal"
	"github.com/mrandycodes/kyber/internal/platform/storage/in_memory"
	"log"
	"net/http"
)

const httpAddr = ":8080"
var repository routes.RoutesRepository

type addRequest struct {
	Route string `json:"route"`
}

func main() {
	fmt.Println("Starting server on port", httpAddr)

	repository = in_memory.NewRoutesRepository()

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/add", addRouteHandler)

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
	log.Println(addReq.Route)

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
