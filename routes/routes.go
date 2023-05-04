package routes

import (
	"jwt_server/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	//app.Static("/static", "./static")

	app.Get("/", controllers.Root)

	app.Get("/api/ping", controllers.Ping)

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
