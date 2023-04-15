package keycloak

import "github.com/Nerzal/gocloak/v13"

type KeycloakRepository interface {
	Authenticate(username string, password string) (*gocloak.JWT, error)
	CreateUser(user gocloak.User) (*string, error)
	RefreshUserToken(RefreshToken string) (*gocloak.JWT, error)
}
