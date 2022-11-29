package contract

import "github.com/golang-jwt/jwt"

type IJWT interface {
	Generate(data jwt.MapClaims) (string, error)
	Validate(token string) (map[string]interface{}, error)
}
