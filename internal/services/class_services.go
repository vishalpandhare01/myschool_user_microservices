package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
)

func AddClassServices(body *model.ClassAndStandrd) interface{} {
	response, err := repository.AddNewClassRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    201,
		Message: "Class Added Successfully",
		Data:    response,
	}
}

func GetClassBySchoolIdServices(schoolId string) interface{} {
	response, err := repository.GetClassBySchoolIdRepository(schoolId)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "Success",
		Data:    response,
	}
}
