package repeat

import (
	"bytes"
	routes "github.com/mrandycodes/kyber/internal"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type repetitionResponse struct {
	Path     string `json:"path"`
	Status   string `json:"status"`
	Response string `json:"response"`
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
