package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/unownone/go-chat/app/user"
	"github.com/unownone/go-chat/middleware"
)

// Base Route = /api/v1/auth
func Auth(base string, app fiber.Router) {
	// V1 Links
	v1 := app.Group("/v1", GetNextMiddleWare)
	v1.Post(base+"/register", auth.Signup)
	v1.Post(base+"/login", auth.Login)
	v1.Post(base+"/refresh", auth.RefreshToken)
	v1.Get(base+"/current-user", middleware.VerifyJwtWithClaim(auth.CurrentUser))
	v1.Get(base+"/user", auth.GetUserByEmail)
}
