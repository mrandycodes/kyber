package in_memory

import (
	routes "github.com/mrandycodes/kyber/internal"
)

type RoutesRepository struct {
	routes []routes.Route
}

func NewRoutesRepository() *RoutesRepository {
	return &RoutesRepository{}
}

func (r *RoutesRepository) Add(route routes.Route) error {
	if r.exists(route) {
		return nil
	}

	r.routes = append(r.routes, route)

	return nil
}

func (r *RoutesRepository) Delete(route routes.Route) error {
	temp := []routes.Route{}
	for _, value := range r.routes {
		if route.Value() != value.Value() {
			temp = append(temp, value)
		}
	}

	r.routes = temp

	return nil
}

func (r RoutesRepository) List() []routes.Route {
	return r.routes
}

func (r RoutesRepository) exists(otherRoute routes.Route) bool {
	for _, route := range r.List() {
		if route.Value() == otherRoute.Value() {
			return true
		}
	}

	return false
}
