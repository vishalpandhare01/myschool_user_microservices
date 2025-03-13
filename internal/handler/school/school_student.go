package school

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

func AddSchoolStudentHandler(c *fiber.Ctx) error {

	var body model.Student
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userId is not a string")
	}

	body.RegisterNumber = int64(body.RegisterNumber)
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	body.SchoolID = schoolId
	response := services.AddNewStudentServices(&body)

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

// GetAllStudentServices
func GetAllSchoolStudentHandler(c *fiber.Ctx) error {
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
	schoolId, ok := c.Locals("userId").(string)
	if !ok {
		fmt.Println("userId is not a string")
	}
	mobileNumber := c.Query("mobileNumber")
	registerNumber := c.Query("registerNumber")
	email := c.Query("email")
	classID := c.Query("class_id")
	fName := c.Query("fName")
	lName := c.Query("lName")

	response := services.GetAllStudentServices(pageStr, limitStr, schoolId, mobileNumber, registerNumber, email, classID, fName, lName)

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

// get student by id
func GetStudentByIdHandler(c *fiber.Ctx) error {
	studentId := c.Params("studentId")
	response := services.GetStudentByIdServices(studentId)

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
