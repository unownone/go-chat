package auth

import (
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/unownone/go-chat/app/response"
)

var (
	access    = jwt.NewNumericDate(time.Now().Add(time.Hour * 1))
	refresh   = jwt.NewNumericDate(time.Now().Add(time.Hour * 120))
	SecretKey = []byte(os.Getenv("JWT_KEY"))
)

func GetJwtToken(user string) (string, string, int) {
	a := 1
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user,
		ExpiresAt: access,
	})
	claims2 := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    user,
		Subject:   "Refresh",
		ExpiresAt: refresh,
	})
	token, err := claims.SignedString([]byte(SecretKey))
	reftoken, err1 := claims2.SignedString([]byte(SecretKey))
	if err != nil && err1 != nil {
		a = 0
	}
	return token, reftoken, a
}

func RefreshToken(c *fiber.Ctx) error {
	tokens := new(response.TokenData)
	claims := &jwt.RegisteredClaims{}
	err := c.BodyParser(tokens)
	if err != nil {
		return c.JSON(response.Error{
			Message: err.Error(),
		})
	}
	_, err = jwt.ParseWithClaims(tokens.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return c.JSON(response.Error{
			Message: err.Error(),
		})
	}
	if claims.Subject != "Refresh" {
		return c.JSON(response.Error{
			Message: "Invalid Token.Not a refresh token.",
		})
	}
	token, refresh, err1 := GetJwtToken(claims.Issuer)
	if err1 != 1 {
		return c.JSON(response.Error{
			Message: "Error while creating token",
		})
	}

	tokens.AccessToken = token
	tokens.RefreshToken = refresh

	return c.JSON(tokens)
}

func VerifyJwtSess(token string) (string, bool) {
	claims := &jwt.RegisteredClaims{}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return "", false
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		return "", false
	}
	return claims.Issuer, true
}
