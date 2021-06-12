package main

import (
	"log"

	"github.com/mrandycodes/kyber/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
