package api

import (
	"github.com/gofiber/fiber/v2"
	"password_warehouse.com/cmd/api/handlers"
)

func New(user handlers.UserHandlerClient) *fiber.App {
	app := fiber.New()

	app.Post("/api/v1/user/create", user.CreateUser)

	//api := app.Group("/api", middleware)

	return app
}
