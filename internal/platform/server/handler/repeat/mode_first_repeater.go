package repeat

import (
	"bytes"
	"encoding/json"
	routes "github.com/mrandycodes/kyber/internal"
	"io/ioutil"
	"net/http"
	"strings"
)

func ModeFirstRepetitionHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		routesList := repository.List()
		path := strings.TrimPrefix(req.RequestURI, "/api")
		method := req.Method
		bodyInBytes, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyInBytes))
		headers := req.Header.Clone()

		channel := make(chan repetitionResponse)

		for _, route := range routesList {
			go func(channel chan repetitionResponse, route routes.Route) {
				channel <- doRequest(method, route, path, bodyInBytes, headers)
			}(channel, route)
		}

		firstResponse := <- channel

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(firstResponse)
	}
}

