package keycloak

import (
	keycloakRepository "github.com/mviniciusgc/authorization/src/repositories/keycloak"
)

type GoCloakClientController struct {
	keycloakRepository keycloakRepository.KeycloakRepository
}

func InitializeKeycloakController(kcr keycloakRepository.KeycloakRepository) KeycloakController {
	return &GoCloakClientController{
		keycloakRepository: kcr,
	}
}
