package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
)

func AddNewStudentServices(body *model.Student) interface{} {
	if body.ClassID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Class Id required",
		}
	}
	_, err := repository.GetClassByIdRepository(body.ClassID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Class Not exist",
			}
		}
	}

	response, err := repository.AddNewStudentRepository(body)
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

func GetAllStudentServices(pageStr string, limitStr string, schoolId string) interface{} {
	response, err := repository.GetAllStudentRepository(pageStr, limitStr, schoolId)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "success",
		Data:    response,
	}
}
