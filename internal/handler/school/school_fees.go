package school

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

// add fees type handler
func AddFeesTypeHandler(c *fiber.Ctx) error {
	var body *model.FeeType

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

	response := services.AddFeeTypesServices(body)

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

// get fees type handler
func GetFeesTypesHandler(c *fiber.Ctx) error {
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	response := services.GetFeeTypesServices(schoolId)

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

// add or update fee structure types
func AddFeeStructureTypesHandler(c *fiber.Ctx) error {
	var body *model.FeesStructure

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

	response := services.AddFeeStructureServices(body)

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

// get fees type handler
func GetFeeStructuresHandler(c *fiber.Ctx) error {
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	response := services.GetFeeStructureServices(schoolId)

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

// add student fee
func AddStudentFeesHandler(c *fiber.Ctx) error {
	var body *model.StudentFees

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

	response := services.AddStudentFeesServices(body)

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

// get fees type handler
func GetStudentFeesHandler(c *fiber.Ctx) error {
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}
	userid := c.Query("userid")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	acadmicYear := c.Query("acadmicYear")
	status := c.Query("status")

	response := services.GetStudentFeesServices(pageStr, limitStr, userid, schoolId, acadmicYear, status)

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
