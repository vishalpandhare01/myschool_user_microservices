package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/validation"
)

// add student in school
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

	if body.Email == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: email required ",
		}
	}

	if body.MobileNumber == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: mobile number required ",
		}
	}

	if body.FatherName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: first name required ",
		}
	}

	if body.LastName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: first name required ",
		}
	}

	if body.ParentsMobileNumber == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: perent mobile number required ",
		}
	}

	if body.MotherName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: mothere name required ",
		}
	}

	if body.PlaceOfBirth == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "place of birth required ",
		}
	}

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

	//get school by id
	schoolData, err := repository.GetSchoolByIdRepository(body.SchoolID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Class Not exist",
			}
		}
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	//first register student
	registerResponse, err := repository.CreateUserRepository(
		repository.UserBody{
			Image:        body.Image,
			FirstName:    body.FirstName,
			LastName:     body.LastName,
			SchoolName:   schoolData.SchoolName,
			Email:        body.Email,
			MobileNumber: body.MobileNumber,
			Role:         "student",
			Address:      body.Address,
		},
	)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: "Error: " + err.Error(),
		}
	}

	body.UserID = registerResponse.ID
	//add student to school

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

// get adll student by schoolid (userid)
func GetAllStudentServices(
	pageStr string,
	limitStr string,
	schoolId string,
	mobileNumber string,
	registerNumber string,
	email string,
	classID string,
	fName string,
	lName string,
) interface{} {
	response, err := repository.GetAllStudentRepository(pageStr, limitStr, schoolId, mobileNumber, registerNumber, email, classID, fName, lName)
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

// get student by id
func GetStudentByIdServices(studentId string) interface{} {
	response, err := repository.GetStudentRepository(studentId)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "student: " + err.Error(),
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

// UpdateStudentRepository
func UpdateStudentDetialsServices(body *model.Student) interface{} {

	//get student first
	studentData, err := repository.GetStudentRepository(body.ID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Studen: " + err.Error(),
			}
		}
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	if body.Email != "" {
		if validation.CheckEmailExist(body.Email) {
			return utils.ErrorResponse{
				Code:    400,
				Message: "Email Aleready in use",
			}
		}

	}
	if body.MobileNumber != "" {
		if validation.CheckEmailExist(body.Email) {
			return utils.ErrorResponse{
				Code:    400,
				Message: "Mobile number Aleready in use",
			}
		}

	}

	print("studentData: ", studentData)
	response, err := repository.UpdateStudentRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "student Updated successfully",
		Data:    response,
	}
}
