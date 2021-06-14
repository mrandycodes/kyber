package repeat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
func doRequest(method string, route routes.Route, path string, bodyInBytes []byte, headers http.Header) repetitionResponse {
	body := strings.NewReader(string(bodyInBytes))
	newRequest, _ := http.NewRequest(method, route.Value()+path, body)
	for key, value := range headers {
		newRequest.Header.Add(key, value[len(value)-1])
	}
	response, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Error: ", err)
			}
		}()
		panic(err.Error())
	}
	newResponseBody, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	response.Body = ioutil.NopCloser(bytes.NewBuffer(newResponseBody))
	return repetitionResponse{
		Path:     route.Value() + path,
		Status:   response.Status,
		Response: string(newResponseBody),
	}
}
