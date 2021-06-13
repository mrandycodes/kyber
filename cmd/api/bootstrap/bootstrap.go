package bootstrap

import (
	"flag"

	"github.com/mrandycodes/kyber/internal/platform/server"
	"github.com/mrandycodes/kyber/internal/platform/storage/in_memory"
)

const (
	_defaultHost = "localhost"
	_defaultPort = 8080
)

func Run() error {
	repository := in_memory.NewRoutesRepository()

	var (
		host = flag.String("host", _defaultHost, "if you run your server on a different host")
		port = flag.Int("port", _defaultPort, "port to run the server")
	)

	flag.Parse()

	srv := server.New(*host, *port, repository)
	return srv.Run()
}
