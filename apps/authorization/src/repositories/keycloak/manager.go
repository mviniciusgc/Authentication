package keycloak

import (
	"fmt"

	"github.com/Nerzal/gocloak/v11"
)

func (s *GoCloakClientRepository) CreateUserTEST(user *gocloak.User) (string, error) {
	fmt.Println("bateu aqui no repository")
	//kc.RefreshToken()
	// ctx := context.Background()
	// credentials := gocloak.CredentialRepresentation{
	// 	Type:      gocloak.StringP("password"),
	// 	Value:     gocloak.StringP("123456"),
	// 	Temporary: gocloak.BoolP(false),
	// }

	// newUser := gocloak.User{
	// 	FirstName:   gocloak.StringP("marcos"),
	// 	LastName:    gocloak.StringP("vinicius"),
	// 	Email:       gocloak.StringP("marcos@gmail.com"),
	// 	Enabled:     gocloak.BoolP(true),
	// 	Username:    gocloak.StringP("mvinicius"),
	// 	Credentials: &[]gocloak.CredentialRepresentation{credentials},
	// }

	//return kc.Client.CreateUser(ctx, kc.Token.AccessToken, kc.Realm, newUser)
	return "ok", nil
}
