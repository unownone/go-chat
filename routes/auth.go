package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/unownone/go-chat/app/user"
	"github.com/unownone/go-chat/middleware"
)

func Auth(base string, app *fiber.App) {

	app.Post(base+"/register", auth.Signup)
	app.Get(base+"/current-user", middleware.VerifyJwt(auth.CurrentUser))
}
