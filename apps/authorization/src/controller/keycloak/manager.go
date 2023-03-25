package keycloak

import (
	"fmt"
)

func (se *GoCloakClient) CreateUser(user User) (string, error) {
	fmt.Println("bateu no controller")
	//se.CreateUserTEST(nil)
	fmt.Println("bateu aqui ", user)

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
