package repeat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	routes "github.com/mrandycodes/kyber/internal"
)

func RepetitionHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")
		routesList := repository.List()
		numberOfRoutes := len(routesList)
		var responses = make([]repetitionResponse, numberOfRoutes)
		path := strings.TrimPrefix(req.RequestURI, "/api")
		method := req.Method
		bodyInBytes, _ := ioutil.ReadAll(req.Body)
		req.Body.Close()
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyInBytes))
		headers := req.Header.Clone()
		wg := &sync.WaitGroup{}
		wg.Add(numberOfRoutes)
		for i, route := range routesList {
			go func(wg *sync.WaitGroup, i int, route routes.Route) {
				responses[i] = doRequest(method, route, path, bodyInBytes, headers)
				wg.Done()
			}(wg, i, route)
		}
		wg.Wait()
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(responses)
	}
}
