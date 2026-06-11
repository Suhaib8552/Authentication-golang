package jwt_auth

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(tokenString string) (*Claims, error) {

	claims := &Claims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {

			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, errors.New("unexpected signing method")
			}

			return secret_jwt, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	fmt.Printf("Parsed Claims: %+v\n", claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
