package in_memory

import (
	routes "github.com/mrandycodes/kyber/internal"
)

type RoutesRepository struct {
	routes []routes.Route
}