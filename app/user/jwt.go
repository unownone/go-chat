package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SecretKey = os.Getenv("JWT_KEY")

func GetJwtToken(user string) (string, error) {
	t := jwt.NewNumericDate(time.Now().Add(time.Hour * 1))

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user,
		ExpiresAt: t,
	})
	token, err := claims.SignedString([]byte(SecretKey))
	return token, err
}
