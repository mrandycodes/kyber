package repeat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

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

		responses := []repetitionResponse{}
		client := &http.Client{}
		path := strings.TrimPrefix(req.RequestURI, "/api")
		method := req.Method
		body := req.Body
		headers := req.Header.Clone()

		for _, route := range repository.List() {
			newRequest, _ := http.NewRequest(method, route.Value()+path, body)
			for key, value := range headers {
				newRequest.Header.Add(key, value[len(value)-1])
			}

			response, _ := client.Do(newRequest)

			body, _ := ioutil.ReadAll(response.Body)

			responses = append(responses, repetitionResponse{
				Path:     route.Value() + path,
				Status:   response.Status,
				Response: string(body),
			})
		}

		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(responses)
	}
}
