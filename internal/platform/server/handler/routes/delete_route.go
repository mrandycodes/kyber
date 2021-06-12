package routes

import (
	"encoding/json"
	"net/http"

	routes "github.com/mrandycodes/kyber/internal"
)

type deleteRouteRequest struct {
	Route string `json:"route"`
}

func DeleteRouteHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
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
}
