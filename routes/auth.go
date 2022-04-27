package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/unownone/gweeter/app/user"
	"github.com/unownone/gweeter/middleware"
)

func Auth(base string, app *fiber.App) {

	app.Post(base+"/register", auth.Signup)
	app.Get(base+"/current-user", middleware.VerifyJwt(auth.CurrentUser))
}
