package keycloak

type GoCloakClient struct {
	ClientID     string
	ClientSecret string
	Pass         string
	Realm        string
	User         string
}

func InitializeKeycloakController() KeycloakController {
	return &GoCloakClient{
		ClientID:     "",
		ClientSecret: "",
		Pass:         "",
		Realm:        "",
		User:         "",
	}
}
