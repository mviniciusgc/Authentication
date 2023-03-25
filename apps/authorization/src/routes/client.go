package routes

import (
	"github.com/go-chi/chi/v5"
	keycloakController "github.com/mviniciusgc/authorization/src/controller/keycloak"
	keycloakRepository "github.com/mviniciusgc/authorization/src/repositories/keycloak"
)

type HandlerServices struct {
	KeycloakController keycloakController.KeycloakController
	KeycloakRepository keycloakRepository.KeycloakRepository
	Route              *chi.Mux
}

func (se HandlerServices) CreateRouterServices() *HandlerServices {
	a := keycloakController.InitializeKeycloakController()
	d := &HandlerServices{
		KeycloakController: a,
	}
	r := Handlers(d)
	se.Route = r
	return &se
}

// func initializeKeycloakController() keycloakController.KeycloakController {
// 	return &keycloakController.GoCloakClient{
// 		ClientID:     "",
// 		ClientSecret: "",
// 		Pass:         "",
// 		Realm:        "",
// 		User:         "",
// 	}
// }
