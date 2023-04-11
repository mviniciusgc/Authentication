package server

import (
	"fmt"
	"log"
	"net/http"
)

func (s Services) InitializeServer() {
	r, err := s.HandlerServer.CreateRouterServices()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to initialize services: %w", err))
	}

	err = http.ListenAndServe(":8082", r.Route)
	if err != nil {
		log.Fatal("Error initialize the server")
	}

}
