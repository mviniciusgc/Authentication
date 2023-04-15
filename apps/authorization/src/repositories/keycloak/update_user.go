package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/utils/middleware"
)

func (s *GoCloakClientRepository) UpdateUser(user gocloak.User) error {
	s.RefreshToken(s.MainRealm)
	ctx := context.Background()

	err := s.Client.UpdateUser(ctx, s.Token.AccessToken, s.Realm, user)
	err = middleware.VerifyErrors(err, "UpdateUser")
	if err != nil {
		return err
	}

	return nil
}
