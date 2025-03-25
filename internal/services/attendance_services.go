package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/validation"
)

// create attaendance by staff teacher
func CreateAttendanceServices(body *model.Attendance) interface{} {
	if body.ClassID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "ClassID required",
		}
	}

	if !repository.CheckClassExistbyClassIdRepository(body.ClassID, body.SchoolID) {
		return utils.ErrorResponse{
			Code:    404,
			Message: "Class Not found",
		}
	}

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required",
		}
	}

	if body.TeacherID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "TeacherID required",
		}
	}

	if !repository.CheckStaffExistStaffRepository(body.TeacherID, body.SchoolID) {
		return utils.ErrorResponse{
			Code:    404,
			Message: "TeacherID not exist",
		}
	}

	if body.StudentID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "StudentID required",
		}
	}

	_, err := repository.GetStudentRepository(body.StudentID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "StudentID not exist",
			}
		}
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	if body.Subject == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Subject required",
		}
	}

	if body.Date == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Date required DD-MM-YYYY",
		}
	}

	if !validation.ValidateDate(body.Date) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Date Format should DD-MM-YYYY",
		}
	}

	//'present','absent,'leave'
	if body.Status != "present" && body.Status != "absent" && body.Status != "leave" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Status required 'present','absent,'leave'",
		}
	}

	response, err := repository.CreateAttendanceRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    201,
		Message: "Attendance Marked successfully",
		Data:    response,
	}
}

// Update attaendance by staff teacher
func UpdateAttendanceServices(body *model.Attendance) interface{} {

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required",
		}
	}

	if body.ID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "attandance ID required",
		}
	}

	if body.ClassID != "" {
		if !repository.CheckClassExistbyClassIdRepository(body.ClassID, body.SchoolID) {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Class Not found",
			}
		}
	}

	if body.TeacherID != "" {
		if !repository.CheckStaffExistStaffRepository(body.TeacherID, body.SchoolID) {
			return utils.ErrorResponse{
				Code:    404,
				Message: "TeacherID not exist",
			}
		}
	}

	if body.StudentID != "" {
		_, err := repository.GetStudentRepository(body.StudentID)
		if err != nil {
			if err.Error() == "record not found" {
				return utils.ErrorResponse{
					Code:    404,
					Message: "StudentID not exist",
				}
			}
			return utils.ErrorResponse{
				Code:    500,
				Message: err.Error(),
			}
		}

	}

	if body.Date != "" {
		if !validation.ValidateDate(body.Date) {
			return utils.ErrorResponse{
				Code:    400,
				Message: "Date Format should DD-MM-YYYY",
			}
		}
	}

	if body.Status != "" {
		//'present','absent,'leave'
		if body.Status != "present" && body.Status != "absent" && body.Status != "leave" {
			return utils.ErrorResponse{
				Code:    400,
				Message: "Status required 'present','absent,'leave'",
			}
		}
	}

	response, err := repository.CreateAttendanceRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "Attendance Mark Updated successfully",
		Data:    response,
	}
}

// get attendance
func GetAttaendanceRepository(
	date string,
	classId string,
	schoolId string,
	subject string,
	teacherId string,
	studentId string,
) interface{} {
	repository, err := repository.GetAttendanceRepository(date, classId, schoolId, subject, teacherId, studentId)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "success",
		Data:    repository,
	}
}
