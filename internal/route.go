package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/handler/admin"
	"github.com/vishalpandhare01/internal/handler/user"
)

func Server(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "server running successfully",
	})
}

func RouteSetUp(app *fiber.App) {
	var userRoutes = app.Group("/api/v1/user")
	userRoutes.Post("/register", user.CreateUserHandler) //access for schools and admin
	userRoutes.Post("/sendOtp", user.SendOtp)
	userRoutes.Post("/veryfyOtp", user.VeryfyOtp)

	var adminRoutes = app.Group("/api/v1/admin")
	adminRoutes.Get("/schools", admin.GetSchoolHandler)
}
