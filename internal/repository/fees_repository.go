package repository

import (
	"fmt"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
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
	if err := initializer.DB.Create(&body).Error; err != nil {
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
func GetStudentFeesRepository(UserId string, schoolID string, acadmicYear string) (*[]model.StudentFees, error) {
	var StudentFees *[]model.StudentFees
	query := initializer.DB.Where("user_id = ? AND school_id = ?", UserId, schoolID)
	if acadmicYear != "" {
		query = initializer.DB.Where("user_id = ? AND school_id = ? AND academic_year = ?", UserId, schoolID, acadmicYear)
	}
	if err := query.
		First(&StudentFees).
		Error; err != nil {
		return nil, err
	}
	return StudentFees, nil
}
