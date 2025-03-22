package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/validation"
)

// add feetypes by school id (userid)
func AddFeeTypesServices(body *model.FeeType) interface{} {
	if body.FeeName == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Fees Name required ",
		}
	}

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required ",
		}
	}

	response, err := repository.AddFeesTypeRepository(body)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Fees: " + err.Error(),
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

// get feetypes by school id (userid)
func GetFeeTypesServices(schoolId string) interface{} {
	response, err := repository.GetFeeTypeListRepository(schoolId)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Fees: " + err.Error(),
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

// add feeStructure by school id (userid)
func AddFeeStructureServices(body *model.FeesStructure) interface{} {

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required ",
		}
	}

	if body.FeeTypeID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "FeeTypeID required",
		}
	}

	if !repository.CheckFeeTypeListRepository(body.SchoolID, body.FeeTypeID) {
		return utils.ErrorResponse{
			Code:    404,
			Message: "FeeTypeID not exist ",
		}
	}

	if body.AcademicYear == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "FAcademicYear required (YYYY-YYYY)",
		}
	}

	if !validation.ValidateYearRange(body.AcademicYear) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Invalid year format (YYYY-YYYY)",
		}
	}

	if body.Amount == 0 {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Amount required",
		}
	}

	response, err := repository.AddFeesStructureRepository(body)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Fees: " + err.Error(),
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

// get feeStructure by school id (userid)
func GetFeeStructureServices(schoolId string) interface{} {
	response, err := repository.GetFeesStructureListRepository(schoolId)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Fees: " + err.Error(),
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

// add student fees
func AddStudentFeesServices(body *model.StudentFees) interface{} {
	if body.AcademicYear == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "FAcademicYear required (YYYY-YYYY)",
		}
	}

	if !validation.ValidateYearRange(body.AcademicYear) {
		return utils.ErrorResponse{
			Code:    400,
			Message: "Invalid year format (YYYY-YYYY)",
		}
	}

	if body.UserID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "UserID required ",
		}
	}

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "SchoolID required ",
		}
	}

	//'paid','pending','partial'
	if body.Status != "paid" && body.Status != "pending" && body.Status != "partial" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "status should be 'paid','pending','partial'",
		}
	}

	if body.PaidAmount == 0 {
		return utils.ErrorResponse{
			Code:    400,
			Message: "PaidAmount required ",
		}
	}

	if body.RemainingAmount == 0 {
		return utils.ErrorResponse{
			Code:    400,
			Message: "RemainingAmount required ",
		}
	}

	if body.TotalAmount != 0 {
		return utils.ErrorResponse{
			Code:    400,
			Message: "TotalAmount required ",
		}
	}

	if body.FeesDetails == "" {
		return utils.ErrorResponse{
			Code:    400,
			Message: "FeesDetails required ",
		}
	}

	response, err := repository.AddFeesStudentFeesRepository(body)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Fees: " + err.Error(),
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

// get student fees by user id and schoolid or acadmic year
func GetStudentFeesServices(userid string, schoolId string, acdmicyear string) interface{} {
	response, err := repository.GetStudentFeesRepository(userid, schoolId, acdmicyear)
	if err != nil {
		if err.Error() == "record not found" {
			return utils.ErrorResponse{
				Code:    404,
				Message: "Fees: " + err.Error(),
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
