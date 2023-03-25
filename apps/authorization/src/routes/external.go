package routes

import (
	"fmt"
	"net/http"
)

func createUser(s *HandlerServices) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Printf("%+v ", s.KeycloakManager)
		s.KeycloakManager.CreateUserTEST("dd")
	})
}
