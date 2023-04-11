package routes

import (
	"github.com/go-chi/chi/v5"
	keycloakController "github.com/mviniciusgc/authorization/src/controller/keycloak"
	keycloakRepository "github.com/mviniciusgc/authorization/src/repositories/keycloak"
)

type HandlerServices struct {
	KeycloakController keycloakController.KeycloakController
	Route              *chi.Mux
}

func (se HandlerServices) CreateRouterServices() (*HandlerServices, error) {
	krr, err := keycloakRepository.InitializeKeycloakRepository()
	if err != nil {
		return nil, err
	}
	krc := keycloakController.InitializeKeycloakController(krr)
	d := &HandlerServices{
		KeycloakController: krc,
	}
	r := Handlers(d)
	se.Route = r
	return &se, nil
}
