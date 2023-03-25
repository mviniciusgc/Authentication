package main

import (
	"authorization/src/server"
)

type mainService struct {
	Server server.Services
}

func main() {
	var s mainService
	s.Server.InitializeServer()
}
