package school

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/internal/services"
	"github.com/vishalpandhare01/internal/utils"
)

// get Attendance handler
// date string, classId string, schoolId string, subject string, teacherId string, studentId string
func GetAttendanceHandler(c *fiber.Ctx) error {
	date := c.Query("date")
	classId := c.Query("classId")
	schoolId := c.Query("schoolId")
	subject := c.Query("subject")
	teacherId := c.Query("teacherId")
	studentId := c.Query("studentId")

	response := services.GetAttaendanceRepository(date, classId, schoolId, subject, teacherId, studentId)
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
