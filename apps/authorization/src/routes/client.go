package routes

import (
	"github.com/go-chi/chi/v5"
	keycloakController "github.com/mviniciusgc/authorization/src/controller/keycloak"
	keycloakRepository "github.com/mviniciusgc/authorization/src/repositories/keycloak"
	"github.com/mviniciusgc/authorization/src/utils/middleware"
)

type HandlerServices struct {
	KeycloakController keycloakController.KeycloakController
	Route              *chi.Mux
	middleware         middleware.Middleware
}

func (se HandlerServices) CreateRouterServices() (*HandlerServices, error) {
	krr, err := keycloakRepository.InitializeKeycloakRepository()
	if err != nil {
		return nil, err
	}

	krc := keycloakController.InitializeKeycloakController(krr)
	mw := middleware.InitializeMiddleware()
	d := &HandlerServices{
		KeycloakController: krc,
		middleware:         mw,
	}
	r := Handlers(d)
	se.Route = r
	return &se, nil
}
