package middleware

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/mviniciusgc/authorization/src/utils/errors"
)

// func ValidateToken(tokenString string) (jwt.MapClaims, error) {

// 	secretKey, err := getPublicKey()
// 	if err != nil {
// 		return nil, err
// 	}
// 	tokenString, err = getTokenFormated(tokenString)
// 	//fmt.Printf("new token sadasd %+v\n", tokenString)
// 	if err != nil {
// 		fmt.Println("erro 1.1")
// 		return nil, err
// 	}
// 	// Parse the JWT token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Validate the algorithm used to sign the token
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			fmt.Println("erro 1.2")
// 			fmt.Printf("método de assinatura inesperado: %+v\n", token.Header["alg"])
// 			return nil, errors.NewError(&errors.Error{Message: "Unexpected signature method", Code: errors.EUNAUTHORIZED})
// 		}
// 		// Return the secret key used to sign the token
// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// 	// Validate the token's claims
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// Verify that the token has not expired
// 		exp := int64(claims["exp"].(float64))
// 		if exp < time.Now().Unix() {
// 			fmt.Println("erro 1")
// 			return nil, errors.NewError(&errors.Error{Op: "ValidateToken", Message: "Token has expired", Code: errors.EUNAUTHORIZED})
// 		}
// 		// Verify that the token was issued by a trusted issuer
// 		if claims["iss"] != "my_issuer" {
// 			fmt.Println("erro 2")
// 			return nil, errors.NewError(&errors.Error{Op: "ValidateToken", Message: "Invalid issuer", Code: errors.EUNAUTHORIZED})
// 		}
// 		// Return the token's claims if all checks pass
// 		return claims, nil
// 	} else {
// 		fmt.Println("erro 3")
// 		return nil, errors.NewError(&errors.Error{Op: "ValidateToken", Message: "Invalid token claims", Code: errors.EUNAUTHORIZED})
// 	}
// }

// func getPublicKey() ([]byte, error) {
// 	publicKey := viper.GetString("PUBLIC_KEY")
// 	publicKeyByte, err := json.Marshal(publicKey)
// 	if err != nil {
// 		return nil, errors.NewError(&errors.Error{Op: "ValidateToken", Message: "Error get publicKey", Code: errors.EUNAUTHORIZED})
// 	}
// 	return publicKeyByte, nil
// }

func getTokenFormated(token string) (string, error) {
	if token == "" || !strings.HasPrefix(token, "Bearer ") {
		return "", errors.NewError(&errors.Error{Op: "ValidateToken", Message: "Missing header", Code: errors.EUNAUTHORIZED})
	}
	tokenString := strings.TrimPrefix(token, "Bearer ")

	return tokenString, nil
}

func (p HandlerServices) ParseToken(tokenString string) error {
	tokenString, err := getTokenFormated(tokenString)
	if err != nil {
		return errors.NewError(&errors.Error{Op: "tokenparser.Auth0TokenParser.parseToken", Message: "roles domain not found in token"})
	}
	_, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		cert, err := p.getPemCert(token)
		if err != nil {
			return nil, errors.NewError(&errors.Error{Op: "tokenparser.Auth0TokenParser.parseToken", Message: "roles domain not found in token"})
		}
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		return result, nil
	})
	if err != nil {
		fmt.Println("dentro do erro de teste ", err)
		return errors.NewError(&errors.Error{Op: "tokenparser.Auth0TokenParser.parseToken", Message: err.Error(), Err: &errors.Error{Code: "409"}})
	}
	return nil
}

func (p HandlerServices) getPemCert(token *jwt.Token) (string, error) {
	var cert string
	for k := range p.jwks.Keys {
		if token.Header["kid"] == p.jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + p.jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}
	if cert == "" {
		return cert, errors.NewError(&errors.Error{Op: "tokenparser.Auth0TokenParser.getPemCert", Message: "unable to find appropriate key"})
	}
	return cert, nil
}
