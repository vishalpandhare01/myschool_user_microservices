package services

import (
	"strings"

	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
)

// create exam time table
func CreateTimeTableServices(body *model.TimeTable) interface{} {
	//todo check class exist
	//todo check staff exist
	//todo sort start time order
	if body.ClassID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "classID required",
		}
	}
	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required",
		}
	}

	if body.Subject == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Subject required",
		}
	}
	if body.StartTime == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "StartTime required",
		}
	}
	if body.EdnTime == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "EdnTime required",
		}
	}

	if body.TeacherID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "TeacherID required",
		}
	}

	body.DayOfWeek = strings.ToUpper(body.DayOfWeek)
	if body.DayOfWeek == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "DayOfWeek required (Sunday , Monday , Tueday , Wednesday ,Friday , seaterday)",
		}
	}

	response, err := repository.CreateTimeTableRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    201,
		Message: "Time Table created successfully",
		Data:    response,
	}
}

// get time table
func GetTimeTableServices(classId string, schoolId string) interface{} {
	response, err := repository.GetTimeTableByClassIdRepository(classId, schoolId)
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

// delete time table by id
func DeleteTimeTableServices(tableId string, schoolId string) interface{} {
	response, err := repository.DeleteTimeTableByIdRepository(schoolId, tableId)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Time Table: " + err.Error(),
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
