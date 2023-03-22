package server

import (
	"fmt"
	"log"
	"net/http"
)

func InitializeServer() {
	fmt.Println("Initialize server in port 8082")
	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal("Error initialize the server")
	}
}
