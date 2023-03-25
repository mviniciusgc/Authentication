package keycloak

type KeycloakController interface {
	CreateUser(user User) (string, error)
}

type User struct {
	FirstName string `json:"first_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
