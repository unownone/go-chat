package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/unownone/go-chat/app/user"
	"github.com/unownone/go-chat/middleware"
)

// Base Route = /api/v1/auth
func Auth(base string, app *fiber.App) {

	app.Post(base+"/register", auth.Signup)
	app.Post(base+"/login", auth.Login)
	app.Post(base+"/refresh", auth.RefreshToken)
	app.Get(base+"/current-user", middleware.VerifyJwtWithClaim(auth.CurrentUser))
}
