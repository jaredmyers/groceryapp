package main

import (
	"flag"

	"github.com/jaredmyers/groceryapp/user_service/go_version/api"
	"github.com/jaredmyers/groceryapp/user_service/go_version/storage"
)

func main() {
	listenAddr := flag.String("listenaddr", ":8001", "server port")
	flag.Parse()

	store := storage.NewMockStorage()

	server := api.NewServer(*listenAddr, store)
	server.Run()
}
