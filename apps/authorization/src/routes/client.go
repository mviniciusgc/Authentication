package routes

import (
	"authorization/src/keycloak"

	"github.com/go-chi/chi/v5"
)

type HandlerServices struct {
	KeycloakManager keycloak.KeycloakManager
	Route           *chi.Mux
}

func (se HandlerServices) CreateRouterServices() *HandlerServices {
	a := initializeAuth0Client()
	d := &HandlerServices{
		KeycloakManager: a,
	}
	r := Handlers(d)
	se.Route = r
	return &se
}

func initializeAuth0Client() keycloak.KeycloakManager {
	return &keycloak.GoCloakClient{
		ClientID:     "",
		ClientSecret: "",
		Pass:         "",
		Realm:        "",
		User:         "",
	}
}
