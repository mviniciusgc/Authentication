package server

import (
	"fmt"
	"log"
	"net/http"
)

func (s Services) InitializeServer() {
	r := s.HandlerServer.CreateRouterServices()

	fmt.Println("Initialize server in port 8082")
	err := http.ListenAndServe(":8082", r.Route)
	if err != nil {
		log.Fatal("Error initialize the server")
	}

}
