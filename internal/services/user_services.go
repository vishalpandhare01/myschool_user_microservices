package services

import (
	"strings"

	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/email"
	"github.com/vishalpandhare01/internal/utils/jwtToken"
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

	if validation.CheckMobileNumberExist(body.MobileNumber) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Mobile Aleready in use",
		}
	}

	// , 'teacher', 'school', 'student'
	if !validation.CheckRoleIsCorrect(body.Role) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Role should be , 'teacher', 'school', 'student'",
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

type CreateOtpBody struct {
	Email string
	Phone string
}

func CreateOtpServices(body CreateOtpBody) interface{} {

	//check email or phone number exist
	var user model.User

	if body.Email != "" {
		response, err := repository.GetUserByEmailOrPhone(body.Email, true)
		if err != nil {
			if err.Error() == "record not found" {
				return utils.ErrorResponse{
					Code:    404,
					Message: "email: " + err.Error(),
				}
			}
			return utils.ErrorResponse{
				Code:    500,
				Message: err.Error(),
			}
		}
		user = *response
	} else {
		response, err := repository.GetUserByEmailOrPhone(body.Phone, false)
		if err != nil {
			if err.Error() == "record not found" {
				return utils.ErrorResponse{
					Code:    404,
					Message: "email: " + err.Error(),
				}
			}
			return utils.ErrorResponse{
				Code:    500,
				Message: err.Error(),
			}
		}
		user = *response
	}

	code := email.Otp_Number_Generate()
	response, err := repository.CreateOtpForUserRepository(repository.OtPBody{
		OtpCode: code,
		UserId:  user.ID,
	})
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	//Todo: send otp by phone and email

	return utils.SuccessResponse{
		Code:    201,
		Message: "Otp Sent Successfully",
		Data:    response.OtpCode, // comment when email and hone otp implimented
	}

}

type VeryfyOtpBody struct {
	Email string
	Phone string
	Otp   int
}

func VeryfyOtpServices(body VeryfyOtpBody) interface{} {
	var user model.User
	if body.Email != "" {
		response, err := repository.GetUserByEmailOrPhone(body.Email, true)
		if err != nil {
			if err.Error() == "record not found" {
				return utils.ErrorResponse{
					Code:    404,
					Message: "email: " + err.Error(),
				}
			}
			return utils.ErrorResponse{
				Code:    500,
				Message: err.Error(),
			}
		}
		user = *response
	} else {
		response, err := repository.GetUserByEmailOrPhone(body.Phone, false)
		if err != nil {
			if err.Error() == "record not found" {
				return utils.ErrorResponse{
					Code:    404,
					Message: "phone: " + err.Error(),
				}
			}
			return utils.ErrorResponse{
				Code:    500,
				Message: err.Error(),
			}
		}
		user = *response
	}
	_, err := repository.VeryfyOtpRepository(body.Otp, user.ID)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	token, err := jwtToken.GenerateToken(user.ID, user.Role)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "Otp veryfied successfully",
		Data:    token,
	}
}
