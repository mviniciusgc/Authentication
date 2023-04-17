package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/repositories/keycloak/utils"
)

func (s *GoCloakClientRepository) RefreshUserToken(RefreshToken string) (*gocloak.JWT, error) {
	s.RefreshToken(s.MainRealm)
	ctx := context.Background()

	token, err := s.Client.RefreshToken(ctx, RefreshToken, s.ClientID, s.ClientSecret, s.Realm)
	err = utils.VerifyErrors(err, "RefreshUserToken")
	if err != nil {
		return nil, err
	}

	return token, nil
}
