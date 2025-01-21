package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

func CreateUserHandler(c *fiber.Ctx) error {
	var body repository.UserBody

	response := services.CreateUserServices(body)

	switch r := response.(type) {
	case utils.ErrorResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
		})
	case utils.SuccessResponse:
		return c.Status(r.Code).JSON(fiber.Map{
			"message": r.Message,
			"data":    r.Data,
		})
	default:
		return c.Status(500).JSON(fiber.Map{
			"message": "Somthing wrong in services",
		})

	}

}
