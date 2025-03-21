package repository

import (
	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

// add fees types
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

// add fees structure
func AddFeesStructureRepository(body *model.FeesStructure) (*model.FeesStructure, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}

	return body, nil
}
