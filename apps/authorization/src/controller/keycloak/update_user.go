package keycloak

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/entity"
)

func (se *GoCloakClientController) UpdateUser(authenticate entity.UserUpdateRequest, userID string) error {
	userUpdate := gocloak.User{ID: &userID}

	if authenticate.Email != nil {
		userUpdate.Email = authenticate.Email
	}

	err := se.keycloakRepository.UpdateUser(userUpdate)
	if err != nil {
		return err
	}

	return nil
}
