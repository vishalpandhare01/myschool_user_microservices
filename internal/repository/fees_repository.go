package repository

import (
	"fmt"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/funcation"
)

// add and update fees types
func AddFeesTypeRepository(body *model.FeeType) (*model.FeeType, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}

	return body, nil
}

// get fees types
func GetFeeTypeListRepository(schoolID string) (*[]model.FeeType, error) {
	var feetype *[]model.FeeType
	if err := initializer.DB.Where("school_id = ?", schoolID).Find(&feetype).Error; err != nil {
		return nil, err
	}
	return feetype, nil
}

// check feetype exist by id
func CheckFeeTypeListRepository(schoolID string, id string) bool {
	var feetype *model.FeeType
	if err := initializer.DB.Where("school_id = ? AND id = ?", schoolID, id).First(&feetype).Error; err != nil {
		fmt.Println("Error in check Feetypes: ", err.Error())
		return false
	}
	return true
}

// add and update fees structure
func AddFeesStructureRepository(body *model.FeesStructure) (*model.FeesStructure, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}

	return body, nil
}

// get fees  fees structure
func GetFeesStructureListRepository(schoolID string) (*[]model.FeesStructure, error) {
	var FeesStructure *[]model.FeesStructure
	if err := initializer.DB.Where("school_id = ?", schoolID).Find(&FeesStructure).Error; err != nil {
		return nil, err
	}
	return FeesStructure, nil
}

// add  StudentFees
func AddFeesStudentFeesRepository(body *model.StudentFees) (*model.StudentFees, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}

	// if err := initializer.
	// 	DB.Where("user_id = ?", body.User.ID).
	// 	First(&body).Error; err != nil {
	// 	return nil, err
	// }

	// for i := range body.FeesDetails {
	// 	body.TotalAmount += body.FeesDetails[i].Amount
	// }

	// if err := initializer.DB.Save(&body).Error; err != nil {
	// 	return nil, err
	// }

	return body, nil
}

// get  StudentFees by user id and school id
func GetStudentFeesRepository(
	pageStr string,
	limitStr string,
	UserId string,
	schoolID string,
	acadmicYear string,
	status string) (interface{}, error) {

	var StudentFees *[]model.StudentFees
	var totalData []model.StudentFees

	offset, limitInt := funcation.Pagination(pageStr, limitStr)
	fmt.Println("limitInt", limitInt, "offset", offset)

	query := initializer.DB.Where("school_id = ?", schoolID)

	if UserId != "" {
		query = query.Where("user_id = ?", UserId)
	}

	if acadmicYear != "" {
		query = query.Where("academic_year = ?", acadmicYear)
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Now query will be built with all the conditions correctly.

	if err := query.Find(&totalData).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	if err := query.
		Limit(limitInt).
		Offset(offset).
		Preload("User").
		Order("id DESC").
		Find(&StudentFees).
		Error; err != nil {
		return nil, err
	}

	responseData := utils.SuccessListResponse{
		Total:   len(totalData),
		Perpage: limitInt,
		Page:    offset,
		Data:    StudentFees,
	}
	return responseData, nil
}

// get  StudentFees by user id and school id
func GetStudentFeesByIDRepository(UserId string, schoolID string, acadmicYear string) (*model.StudentFees, error) {
	var StudentFees *model.StudentFees
	query := initializer.DB.Where("user_id = ? AND school_id = ? AND academic_year = ?", UserId, schoolID, acadmicYear)
	if err := query.
		First(&StudentFees).
		Error; err != nil {
		return nil, err
	}
	return StudentFees, nil
}
