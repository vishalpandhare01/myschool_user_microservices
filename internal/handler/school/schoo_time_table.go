package school

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

// add Time table handler
func CreateTimeTableHandler(c *fiber.Ctx) error {
	var body *model.TimeTable

	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	body.SchoolID = schoolId

	response := services.CreateTimeTableServices(body)

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

// get Time table handler
func GetTimeTableHandler(c *fiber.Ctx) error {
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}
	classId := c.Params("classId")
	response := services.GetTimeTableServices(classId, schoolId)

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

// delete Time table handler
func DeleteTimeTableHandler(c *fiber.Ctx) error {
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}
	tableId := c.Params("tableId")
	response := services.DeleteTimeTableServices(tableId, schoolId)

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
