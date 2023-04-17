package keycloak

import (
	"context"
	"time"

	"github.com/Nerzal/gocloak/v13"
	"github.com/mviniciusgc/authorization/src/utils/errors"
	"github.com/spf13/viper"
)

type GoCloakClientRepository struct {
	Client        *gocloak.GoCloak
	ClientID      string
	MainRealm     string
	ClientSecret  string
	Pass          string
	User          string
	Realm         string
	Domain        string
	TokenExpireAt time.Time
	Token         *gocloak.JWT
}

func InitializeKeycloakRepository() (KeycloakRepository, error) {
	keycloackUser := viper.GetString("KEYCLOAK_USER")
	keycloackPass := viper.GetString("KEYCLOAK_PASS")
	keycloackDomain := viper.GetString("KEYCLOAK_DOMAIN")
	keycloackRealm := viper.GetString("KEYCLOAK_REALM")
	keycloackClientSecret := viper.GetString("KEYCLOAK_CLIENT_SECRET")
	keycloackClientID := viper.GetString("KEYCLOAK_CLIENT_ID")
	keycloackMainRealm := viper.GetString("KEYCLOAK_MAIN_REALM")

	if keycloackUser == "" || keycloackPass == "" ||
		keycloackDomain == "" || keycloackRealm == "" ||
		keycloackClientSecret == "" || keycloackClientID == "" || keycloackMainRealm == "" {
		return nil, errors.NewError(&errors.Error{Op: "InitializeGoCloakClient", Message: "Missing KeyCloak credentials"})
	}

	client := gocloak.NewClient(keycloackDomain, gocloak.SetAuthAdminRealms("admin/realms"), gocloak.SetAuthRealms("/realms"))
	kc := &GoCloakClientRepository{
		ClientID:     keycloackClientID,
		ClientSecret: keycloackClientSecret,
		Pass:         keycloackPass,
		User:         keycloackUser,
		Realm:        keycloackRealm,
		Client:       client,
		MainRealm:    keycloackMainRealm,
		Domain:       keycloackDomain,
	}
	kc.RefreshToken(keycloackMainRealm)

	return kc, nil
}

func (kc *GoCloakClientRepository) RefreshToken(mainRealm string) error {
	if time.Now().Before(kc.TokenExpireAt) {
		return nil
	}
	ctx := context.Background()
	token, _ := kc.Client.LoginAdmin(ctx, kc.User, kc.Pass, mainRealm)

	kc.Token = token
	kc.TokenExpireAt = time.Now().Add(time.Second * time.Duration(token.ExpiresIn))
	return nil
}
