package jwt_auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_jwt = []byte("34567")

type Claims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int, email string) (string, error) {

	claims := Claims{
		UserID: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	fmt.Printf("Claims before signing: %+v\n", claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secret_jwt)
}
