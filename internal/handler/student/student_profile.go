package student

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

func GetLoginUserHandler(c *fiber.Ctx) error {
	userId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}
	response := services.GetStudentByIdServices(userId)

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
