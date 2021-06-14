package repeat

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	routes "github.com/mrandycodes/kyber/internal"
)

type repetitionResponse struct {
	Path     string `json:"path"`
	Status   string `json:"status"`
	Response string `json:"response"`
}

func RepetitionHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		routes := repository.List()
		numberOfRoutes := len(routes)
		var responses = make([]repetitionResponse, numberOfRoutes)
		path := strings.TrimPrefix(req.RequestURI, "/api")
		method := req.Method
		bodyInBytes, _ := ioutil.ReadAll(req.Body)
		body := strings.NewReader(string(bodyInBytes))
		headers := req.Header.Clone()

		wg := &sync.WaitGroup{}
		wg.Add(numberOfRoutes)
		for i, route := range routes {
			go func(wg *sync.WaitGroup, i int, route routes.Route) {
				responses[i] = doRequest(method, route, path, body, headers)
				wg.Done()
			}(wg, i, route)

		}

		wg.Wait()

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(responses)
	}
}

func doRequest(method string, route routes.Route, path string, body io.Reader, headers http.Header) repetitionResponse {
	client := &http.Client{}

	newRequest, _ := http.NewRequest(method, route.Value()+path, body)
	for key, value := range headers {
		newRequest.Header.Add(key, value[len(value)-1])
	}

	response, _ := client.Do(newRequest)

	newResponseBody, _ := ioutil.ReadAll(response.Body)

	return repetitionResponse{
		Path:     route.Value() + path,
		Status:   response.Status,
		Response: string(newResponseBody),
	}
}
