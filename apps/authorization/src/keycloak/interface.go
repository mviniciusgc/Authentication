package keycloak

type KeycloakManager interface {
	CreateUserTEST(user string) (string, error)
}

type User struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
