package staff

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

// GetAllStudentServices
func GetAllSchoolStudentHandler(c *fiber.Ctx) error {
	schoolId := c.Params("schoolId")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")
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
