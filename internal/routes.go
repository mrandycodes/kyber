package internal

import (
	"fmt"
	"net/url"
)

type Route struct {
	value string
}

func NewRoute(value string) (Route, error) {
	parsedUrl, err := url.ParseRequestURI(value)

	if err != nil {
		return Route{}, fmt.Errorf("invalid route: %s", value)
	}

	return Route{value: parsedUrl.String()}, nil
}

func (r Route) Value() string {
	return r.value
}

type RoutesRepository interface {
	Add(route Route) error
	Delete(route Route) error
	List() []Route
}
