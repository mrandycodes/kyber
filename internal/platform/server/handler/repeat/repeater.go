package repeat

import (
	"encoding/json"
	routes "github.com/mrandycodes/kyber/internal"
	"io/ioutil"
	"net/http"
	"strings"
)

type repetitionResponse struct {
	Path string `json:"path"`
	Status string `json:"status"`
	Response string `json:"response"`
}

func RepetitionHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		responses := []repetitionResponse{}
		client := &http.Client{}
		path := strings.TrimPrefix(req.RequestURI, "/api")
		method := req.Method
		body := req.Body

		for _, route := range repository.List() {
			newRequest, _ := http.NewRequest(method, route.Value() + path, body)

			response, _ := client.Do(newRequest)

			body, _ := ioutil.ReadAll(response.Body)

			responses = append(responses, repetitionResponse{
				Path: route.Value() + path,
				Status: response.Status,
				Response: string(body),
			})
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(responses)
	}
}
