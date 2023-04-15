package keycloak

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/entity"
)

func (se *GoCloakClientController) CreateUser(user entity.UserRequest) (*entity.UserResponse, error) {
	credentials := gocloak.CredentialRepresentation{
		Type:      gocloak.StringP("password"),
		Value:     gocloak.StringP(user.Password),
		Temporary: gocloak.BoolP(false),
	}

	newUser := gocloak.User{
		Email:       gocloak.StringP(user.Email),
		Enabled:     gocloak.BoolP(true),
		Username:    gocloak.StringP(user.Username),
		Credentials: &[]gocloak.CredentialRepresentation{credentials},
	}

	userID, err := se.keycloakRepository.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	newUserRespose := entity.UserResponse{ID: *userID}
	return &newUserRespose, nil
}
