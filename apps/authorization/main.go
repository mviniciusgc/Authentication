package main

import (
	"github.com/mviniciusgc/authorization/src/server"
)

type mainService struct {
	Server server.Services
}

func main() {
	var s mainService
	s.Server.InitializeServer()
}
