package keycloak

import (
	"context"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/repositories/keycloak/utils"
)

func (s *GoCloakClientRepository) Authenticate(username string, password string) (*gocloak.JWT, error) {
	s.RefreshToken(s.MainRealm)
	ctx := context.Background()

	token, err := s.Client.Login(ctx, s.ClientID, s.ClientSecret, s.Realm, username, password)
	err = utils.VerifyErrors(err, "Authenticate")
	if err != nil {
		return nil, err
	}

	return token, nil
}
