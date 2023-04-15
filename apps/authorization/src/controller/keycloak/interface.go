package keycloak

import (
	"github.com/mviniciusgc/authorization/src/entity"
)

type KeycloakController interface {
	Authenticate(authenticate entity.AuthenticateRequest) (*entity.TokenResponse, error)
	CreateUser(user entity.UserRequest) (*entity.UserResponse, error)
	RefreshUserToken(RefreshToken entity.RefreshTokenRequest) (*entity.TokenResponse, error)
}

type User struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
