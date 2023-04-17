package middleware

type Middleware interface {
	ParseToken(tokenString string) error
}
