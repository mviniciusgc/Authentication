package keycloak

import "github.com/Nerzal/gocloak/v11"

type KeycloakRepository interface {
	CreateUserTEST(user *gocloak.User) (string, error)
}
