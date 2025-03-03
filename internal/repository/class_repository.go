package repository

import (
	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

// add class
func AddNewClassRepository(body *model.ClassAndStandrd) (*model.ClassAndStandrd, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

func GetClassBySchoolIdRepository(schoolId string) (*[]model.ClassAndStandrd, error) {
	var data []model.ClassAndStandrd
	if err := initializer.DB.Where("school_id = ?", schoolId).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func GetClassByIdRepository(id string) (*model.ClassAndStandrd, error) {
	var data model.ClassAndStandrd
	if err := initializer.DB.Where("school_id = ?", id).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
