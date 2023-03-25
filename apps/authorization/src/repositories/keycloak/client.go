package keycloak

import "github.com/spf13/viper"

type GoCloakClientRepository struct {
	ClientID     string
	ClientSecret string
	Pass         string
	Realm        string
}

func InitializeKeycloakRepository() KeycloakRepository {
	viper.GetString("ENV")
	return &GoCloakClientRepository{
		ClientID:     "",
		ClientSecret: "",
		Pass:         "",
		Realm:        "",
	}
}
