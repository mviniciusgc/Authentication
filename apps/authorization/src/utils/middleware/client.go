package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/mviniciusgc/authorization/src/utils/errors"
)

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type HandlerServices struct {
	jwks *Jwks
}

func InitializeMiddleware() Middleware {
	mw, _ := createRouterServices()
	return &HandlerServices{jwks: mw}
}

func createRouterServices() (*Jwks, error) {
	jwks, err := getCerts()
	if err != nil {
		return nil, err
	}
	return jwks, nil
}

func getCerts() (*Jwks, error) {
	var jwks Jwks
	//resp, err := http.Get(p.keycloakDomain + "/realms/" + p.realm + "/protocol/openid-connect/certs")
	resp, err := http.Get("http://localhost:8080/realms/master/protocol/openid-connect/certs")
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, errors.NewError(&errors.Error{Op: "tokenparser.KeycloakTokenParser.importJwks", Message: "unable to import jwks.json"})
	}

	json.NewDecoder(resp.Body).Decode(&jwks)
	return &jwks, nil
}
