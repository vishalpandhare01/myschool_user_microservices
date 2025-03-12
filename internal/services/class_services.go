package services

import (
	"fmt"
	"strings"

	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
)

func AddClassServices(body *model.ClassAndStandrd) interface{} {

	if body.ClassName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Class Name required",
		}
	}

	if body.DivisionName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Division Name required",
		}
	}

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required",
		}
	}

	body.DivisionName = strings.ToUpper(body.DivisionName)
	body.ClassName = strings.ToUpper(body.ClassName)

	if repository.CheckClassExistRepository(body) {
		fmt.Println("repository.CheckClassExistRepository(body): ", repository.CheckClassExistRepository(body))
		return utils.ErrorResponse{
			Code:    400,
			Message: "Class Name or Division Name already exists",
		}
	}
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
