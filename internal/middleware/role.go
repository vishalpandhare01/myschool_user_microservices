package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func IsStudent(C *fiber.Ctx) error {
	role, ok := C.Locals("userType").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userType is not a string")
	}

	if role == "student" {
		return C.Next()
	}
	return C.Status(401).JSON(fiber.Map{
		"message": "You Are Not Authorized for this operation please login as student",
	})
}

func IsAdmin(C *fiber.Ctx) error {
	role, ok := C.Locals("userType").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userType is not a string", ok)
	}

	if role == "admin" {
		return C.Next()
	}
	return C.Status(401).JSON(fiber.Map{
		"message": "You Are Not Authorized for this operation please login as admin",
	})
}

func IsTeacher(C *fiber.Ctx) error {
	role, ok := C.Locals("userType").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userType is not a string")
	}

	if role == "teacher" {
		return C.Next()
	}
	return C.Status(401).JSON(fiber.Map{
		"message": "You Are Not Authorized for this operation please login as teacher",
	})
}

func IsSchool(C *fiber.Ctx) error {
	role, ok := C.Locals("userType").(string)
	if !ok {
		// Handle the error if the type assertion fails
		fmt.Println("userType is not a string")
	}

	if role == "school" {
		return C.Next()
	}
	return C.Status(401).JSON(fiber.Map{
		"message": "You Are Not Authorized for this operation please login as school",
	})
}
