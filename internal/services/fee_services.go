package services

import (
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/repository"
	"github.com/vishalpandhare01/internal/utils"
)

// add feetypes by school id (userid)
func AddFeeTypesServices(body *model.FeeType) interface{} {
	if body.FeeName == "" {
		return utils.ErrorResponse{
			Code:    404,
			Message: "Fees Name required ",
		}
	}

	if body.SchoolID == "" {
		return utils.ErrorResponse{
			Code:    404,
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
