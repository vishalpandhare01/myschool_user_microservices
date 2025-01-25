package internal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/handler/admin"
	"github.com/vishalpandhare01/internal/handler/school"
	"github.com/vishalpandhare01/internal/handler/user"
	"github.com/vishalpandhare01/internal/middleware"
)

func Server(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"message": "server running successfully",
	})
}

func RouteSetUp(app *fiber.App) {
	// user apis
	var userRoutes = app.Group("/api/v1/user")
	userRoutes.Post("/register", user.CreateUserHandler) //access for schools and admin
	userRoutes.Post("/sendOtp", user.SendOtp)
	userRoutes.Post("/veryfyOtp", user.VeryfyOtp)

	//admin apis
	var adminRoutes = app.Group("/api/v1/admin", middleware.Authentication, middleware.Authorization)
	adminRoutes.Get("/schools", admin.GetSchoolsHandler)
	adminRoutes.Patch("/school/:schoolId", admin.UpdateSchoolByIdHandler)
	adminRoutes.Get("/school/:schoolId", school.GetSchoolByIdHandler)

	//school apis
	var schoolRoutes = app.Group("/api/v1/school", middleware.Authentication)
	schoolRoutes.Get("/", school.GetSchoolByIdHandler)

}
