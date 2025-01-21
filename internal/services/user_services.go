package services

import (
	"strings"

	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/validation"
)

func CreateUserServices(body repository.UserBody) interface{} {
	body.Role = strings.ToLower(body.Role)

	if validation.CheckEmailExist(body.Email) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Email Aleready in use",
		}
	}

	if validation.CheckMobileNumberExist(body.Email) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Mobile Aleready in use",
		}
	}

	// 'admin', 'teacher', 'school', 'student'
	if !validation.CheckRoleIsCorrect(body.Role) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Role should be 'admin', 'teacher', 'school', 'student'",
		}
	}

	response, err := repository.CreateUserRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: "Error: " + err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    201,
		Message: "Success",
		Data:    response,
	}
}
