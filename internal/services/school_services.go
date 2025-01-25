package services

import (
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
)

func GetSchoolsServices(pageStr string, limitStr string, school_name string, isPaid string) interface{} {
	response, err := repository.GetSchoolsRepository(pageStr, limitStr, school_name, isPaid)
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

func UpdateSchoolsServices(schoolId string) interface{} {
	response, err := repository.UpdateSchoolPaymentRepository(schoolId)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "school: " + err.Error(),
			}
		}
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "school updated successfully",
		Data:    response,
	}
}

func GetSchoolByIdServices(schoolId string) interface{} {
	response, err := repository.GetSchoolByIdRepository(schoolId)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "school: " + err.Error(),
			}
		}
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
