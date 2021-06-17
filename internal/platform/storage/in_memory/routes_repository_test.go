package in_memory

import (
	"testing"

	routes "github.com/mrandycodes/kyber/internal"
	"github.com/stretchr/testify/assert"
)

func Test_RoutesRepository_Add_New_Routes(t *testing.T) {
	repository := NewRoutesRepository()
	route, _ := routes.NewRoute("http://disney.com")

	repository.Add(route)
	routeAdded := repository.List()[0]

	assert.Equal(t, route.Value(), routeAdded.Value())
}

func Test_RoutesRepository_Just_Add_One_Time_The_Same_Route(t *testing.T) {
	repository := NewRoutesRepository()
	route, _ := routes.NewRoute("http://marvel.com")

	repository.Add(route)
	repository.Add(route)

	assert.Equal(t, 1, len(repository.List()))
}

func Test_RoutesRepository_Can_Remove_An_Existing_Route(t *testing.T) {
	repository := NewRoutesRepository()
	routeOne, _ := routes.NewRoute("http://fox.com")
	routeTwo, _ := routes.NewRoute("http://marvel.com")

	repository.Add(routeOne)
	repository.Add(routeTwo)
	repository.Delete(routeOne)

	assert.Equal(t, 1, len(repository.List()))
}

func Test_RoutesRepository_Can_List_Multiple_Routes(t *testing.T) {
	repository := NewRoutesRepository()
	routeOne, _ := routes.NewRoute("http://fox.com")
	routeTwo, _ := routes.NewRoute("http://marvel.com")

	repository.Add(routeOne)
	repository.Add(routeTwo)

	routesList := repository.List()

	assert.Equal(t, 2, len(routesList))
	assert.Equal(t, routeOne.Value(), routesList[0].Value())
	assert.Equal(t, routeTwo.Value(), routesList[1].Value())
}
