package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/utils"
)

func CreateAttendanceServices(body *model.Attendance) interface{} {

	return utils.SuccessResponse{
		Code:    201,
		Message: "Attendance created successfully",
	}
}
