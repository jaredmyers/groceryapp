package main

import (
	"flag"

	"github.com/jaredmyers/groceryapp/backend/api"
	"github.com/jaredmyers/groceryapp/backend/storage"
)

func main() {

	// setting up server port
	listenAddr := flag.String("listenaddr", ":8000", "server port")
	flag.Parse()

	// setting up storage (mock for now)
	store := storage.NewMemoryStorage()

	// setting up microservice information, (local for now)
	userService := "http://localhost:8001"
	services := map[string]string{"userService": userService}

	// start API
	server := api.NewServer(*listenAddr, store, services)
	server.Run()
}
