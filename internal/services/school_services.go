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
