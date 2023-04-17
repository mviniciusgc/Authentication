package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/repositories/keycloak/utils"
)

func (s *GoCloakClientRepository) UpdateUser(user gocloak.User) error {
	s.RefreshToken(s.MainRealm)
	ctx := context.Background()

	err := s.Client.UpdateUser(ctx, s.Token.AccessToken, s.Realm, user)
	err = utils.VerifyErrors(err, "UpdateUser")
	if err != nil {
		return err
	}

	return nil
}
