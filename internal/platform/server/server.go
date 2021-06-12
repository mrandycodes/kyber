package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	routes "github.com/mrandycodes/kyber/internal"
	"github.com/mrandycodes/kyber/internal/platform/server/handler/health"
	routes_handlers "github.com/mrandycodes/kyber/internal/platform/server/handler/routes"
)

type Server struct {
	httpAddr string
	engine   *mux.Router

	repository routes.RoutesRepository
}

func New(host string, port uint, repository routes.RoutesRepository) Server {
	srv := Server{
		engine:   mux.NewRouter(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		repository: repository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	log.Fatal(http.ListenAndServe(s.httpAddr, s.engine))

	return nil
}

func (s *Server) registerRoutes() {
	s.engine.HandleFunc("/health", health.HealthHandler())
	s.engine.HandleFunc("/routes", routes_handlers.AddRouteHandler(s.repository)).Methods("POST")
	s.engine.HandleFunc("/routes", routes_handlers.DeleteRouteHandler(s.repository)).Methods("DELETE")
	s.engine.HandleFunc("/routes", routes_handlers.ListRoutesHandler(s.repository)).Methods("GET")
}
