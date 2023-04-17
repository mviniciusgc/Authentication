package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/repositories/keycloak/utils"
)

func (s *GoCloakClientRepository) GetUserId(userID string) (*gocloak.User, error) {
	s.RefreshToken(s.MainRealm)
	ctx := context.Background()

	user, err := s.Client.GetUserByID(ctx, s.Token.AccessToken, s.Realm, userID)
	err = utils.VerifyErrors(err, "GetUserId")
	if err != nil {
		return nil, err
	}

	return user, nil
}
