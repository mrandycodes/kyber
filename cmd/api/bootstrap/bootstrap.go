package bootstrap

import (
	"github.com/mrandycodes/kyber/internal/platform/server"
	"github.com/mrandycodes/kyber/internal/platform/storage/in_memory"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	repository := in_memory.NewRoutesRepository()

	srv := server.New(host, port, repository)
	return srv.Run()
}
