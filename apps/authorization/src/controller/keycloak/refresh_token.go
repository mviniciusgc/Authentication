package keycloak

import (
	"bytes"
	"encoding/json"

	"github.com/mviniciusgc/authorization/src/entity"
)

func (se *GoCloakClientController) RefreshUserToken(RefreshToken entity.RefreshTokenRequest) (*entity.TokenResponse, error) {
	token, err := se.keycloakRepository.RefreshUserToken(RefreshToken.RefreshToken)
	if err != nil {
		return nil, err
	}

	tokenResp := &entity.TokenResponse{}
	jsonBytes, err := json.Marshal(token)
	if err != nil {
		return nil, err
	}

	requestReader := bytes.NewReader(jsonBytes)
	err = json.NewDecoder(requestReader).Decode(tokenResp)
	if err != nil {
		return nil, err
	}
	return tokenResp, nil
}
