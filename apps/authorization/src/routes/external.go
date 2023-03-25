package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mviniciusgc/authorization/src/controller/keycloak"
)

func createUser(s *HandlerServices) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Println("bateu no external")
		fmt.Printf("%+v ", s.KeycloakController)

		user := &keycloak.User{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println("dentro do erro")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Println("antes da request do erro")
		resp, err := s.KeycloakController.CreateUser(*user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		jRole, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusCreated)
		w.Write(jRole)
	})
}
