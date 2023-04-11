package keycloak

import "github.com/Nerzal/gocloak/v13"

type KeycloakRepository interface {
	CreateUser(user gocloak.User) (*string, error)
}
