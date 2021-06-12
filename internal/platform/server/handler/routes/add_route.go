package routes

import (
	"encoding/json"
	"net/http"

	routes "github.com/mrandycodes/kyber/internal"
)

type addRequest struct {
	Route string `json:"route"`
}

func AddRouteHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
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
}
