package repeat

import (
	routes "github.com/mrandycodes/kyber/internal"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func RepetitionHandler(repository routes.RoutesRepository) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		client := &http.Client{}
		path := strings.TrimPrefix(req.RequestURI, "/api")
		log.Println(path)
		method := req.Method
		body := req.Body

		for _, route := range repository.List() {
			log.Println("Request to: " + route.Value() + path)
			newRequest, _ := http.NewRequest(method, route.Value() + path, body)

			response, _ := client.Do(newRequest)

			body, _ := ioutil.ReadAll(response.Body)
			log.Println("Response: " + response.Status + " Body: " + string(body))
		}

		res.WriteHeader(http.StatusOK)
	}
}
