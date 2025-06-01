package main

import (
	"log"
	"net/http"

	"mws-test/internal/api"
	"mws-test/internal/service"
	"mws-test/internal/store"
)

func main() {
	st := store.NewMemoryStore()
	svc := service.NewCatService(st)

	server, err := api.NewServer(svc)
	if err != nil {
		log.Fatalf("create server: %v", err)
	}

	addr := ":8000"
	log.Printf("Cat service listening on %s", addr)

	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatal(err)
	}
}
