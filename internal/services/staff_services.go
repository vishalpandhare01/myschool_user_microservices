package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/validation"
)

// add staff in school
func AddNewStaffServices(body *model.Staff) interface{} {
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
				Message: "ClassID Not" + err.Error(),
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

	if body.LastName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: first name required ",
		}
	}

	if body.Gender != "male" && body.Gender != "female" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Error: Gender required ",
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
				Message: "schoolId " + err.Error(),
			}
		}
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	//first register staff
	registerResponse, err := repository.CreateUserRepository(
		repository.UserBody{
			Image:        body.Image,
			FirstName:    body.FirstName,
			LastName:     body.LastName,
			SchoolName:   schoolData.SchoolName,
			Email:        body.Email,
			MobileNumber: body.MobileNumber,
			Role:         "staff",
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
	//add staff to school

	response, err := repository.AddNewStaffRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: "Error: " + err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    201,
		Message: "staff added successfully",
		Data:    response,
	}
}

// get all staff services
func GetAllStaffServices(
	pageStr string,
	limitStr string,
	schoolId string,
	mobileNumber string,
	email string,
	classID string,
	fName string,
	lName string,
) interface{} {

	if classID != "" {
		if !repository.CheckClassExistbyClassIdRepository(classID, schoolId) {
			return utils.ErrorResponse{
				Code:    404,
				Message: "classId not exist",
			}
		}
	}

	response, err := repository.GetAllStaffRepository(pageStr, limitStr, schoolId, mobileNumber, email, classID, fName, lName)
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

// get staff by id
func GetStaffByIdServices(userID string) interface{} {
	response, err := repository.GetStaffRepository(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "staff: " + err.Error(),
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

// Update staff services
func UpdateStaffDetialsServices(body *model.Staff) interface{} {
	if body.UserID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "staff UserID required",
		}
	}

	//get staff first
	staffdata, err := repository.GetStaffRepository(body.UserID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "staff: " + err.Error(),
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

	if body.ClassID != "" {
		_, err := repository.GetClassByIdRepository(body.ClassID)
		if err != nil {
			if err.Error() == "record not found" {
				return utils.ErrorResponse{
					Code:    400,
					Message: "classId " + err.Error(),
				}
			}
			return utils.ErrorResponse{
				Code:    500,
				Message: err.Error(),
			}
		}
	}

	print("staffdata: ", staffdata)
	response, err := repository.UpdateStaffRepository(body)
	if err != nil {
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "staff Updated successfully",
		Data:    response,
	}
}

// Delete Staff Acount
func DeleteStaffByIdServices(userID string) interface{} {
	response, err := repository.DeleteStaffAcountSchoolRepository(userID)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "staff: " + err.Error(),
			}
		}
		return utils.ErrorResponse{
			Code:    500,
			Message: err.Error(),
		}
	}

	return utils.SuccessResponse{
		Code:    200,
		Message: "staff account delete successfully",
		Data:    response,
	}
}
