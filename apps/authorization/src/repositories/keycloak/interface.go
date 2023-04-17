package keycloak

import (
	"github.com/Nerzal/gocloak/v13"
)

type KeycloakRepository interface {
	Authenticate(username string, password string) (*gocloak.JWT, error)
	CreateUser(user gocloak.User) (*string, error)
	GetUserId(userID string) (*gocloak.User, error)
	RefreshUserToken(RefreshToken string) (*gocloak.JWT, error)
	UpdateUser(user gocloak.User) error
}
