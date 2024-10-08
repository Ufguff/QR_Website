package main

import (
	"log"

	"github.com/ufguff/cmd/api"
)

func main() {
	server := api.NewApiServer(":8080")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
