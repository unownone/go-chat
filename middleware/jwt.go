package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwt_key = []byte(os.Getenv("JWT_KEY"))

type AuthHeaders struct {
	Authorization string `reqHeader:"Authorization"`
}

func VerifyJwt(wrapper func(c *fiber.Ctx) error) func(c *fiber.Ctx) error {
	//Function Wrapper to handle JWT Verification
	return func(c *fiber.Ctx) error {
		headers := new(AuthHeaders)
		err := c.ReqHeaderParser(headers)
		if err != nil {
			return c.Status(401).SendString("Unauthorized") // 401 Unauthorized
		}

		token := strings.Split(headers.Authorization, " ") // Getting the token

		if len(token) != 2 {
			return c.Status(401).SendString("Invalid Token Format.Unauthorized")
		}

		tkn, err := jwt.Parse(token[1], func(token *jwt.Token) (interface{}, error) { // Parsing/verifying
			return jwt_key, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(401).SendString("Unauthorized 1")
			}
			return c.Status(401).SendString("Unauthorized 2")
		}
		if !tkn.Valid {
			return c.Status(401).SendString("Unauthorized 3")
		}
		return wrapper(c)
	}
}

func VerifyJwtWithClaim(wrapper func(c *fiber.Ctx, claims *jwt.RegisteredClaims) error) func(c *fiber.Ctx) error {
	//Function Wrapper to handle JWT Verification with claims and return to function
	return func(c *fiber.Ctx) error {
		headers := new(AuthHeaders)
		err := c.ReqHeaderParser(headers)
		if err != nil {
			return c.Status(401).SendString("Unauthorized")
		}
		claims := &jwt.RegisteredClaims{}
		token := strings.TrimPrefix(headers.Authorization, "Bearer ")
		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.Status(401).SendString("Unauthorized")
			}
			return c.Status(401).SendString("Unauthorized")
		}
		if !tkn.Valid {
			return c.Status(401).SendString("Unauthorized")
		}
		return wrapper(c, claims)
	}
}
