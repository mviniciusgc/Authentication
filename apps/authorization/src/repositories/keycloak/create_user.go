package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/repositories/keycloak/utils"
)

func (s *GoCloakClientRepository) CreateUser(user gocloak.User) (*string, error) {
	s.RefreshToken(s.MainRealm)
	ctx := context.Background()

	userID, err := s.Client.CreateUser(ctx, s.Token.AccessToken, s.Realm, user)
	err = utils.VerifyErrors(err, "CreateUser")
	if err != nil {
		return nil, err
	}

	return &userID, nil
}
