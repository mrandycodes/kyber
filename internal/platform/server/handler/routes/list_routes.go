package routes

import (
	"encoding/json"
	"net/http"

	routes "github.com/mrandycodes/kyber/internal"
)

type routesResponse struct {
	Route string `json:"route"`
}

func ListRoutesHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		response := []routesResponse{}
		routes := repository.List()

		for _, route := range routes {
			response = append(response, routesResponse{route.Value()})
		}

		json.NewEncoder(res).Encode(response)
	}
}
