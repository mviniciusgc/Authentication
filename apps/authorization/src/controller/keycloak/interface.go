package keycloak

import "github.com/mviniciusgc/authorization/src/entity"

type KeycloakController interface {
	CreateUser(user entity.UserRequest) (*entity.UserResponse, error)
}

type User struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
