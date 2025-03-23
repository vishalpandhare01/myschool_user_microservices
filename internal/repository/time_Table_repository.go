package repository

import (
	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

// create time table by school
func CreateTimeTableRepository(body *model.TimeTable) (*model.TimeTable, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}

	return body, nil
}

// get time table by class id and school id
func GetTimeTableByClassIdRepository(classId string, schoolId string) (*[]model.TimeTable, error) {
	var timeTable *[]model.TimeTable
	if err := initializer.DB.
		Where("class_id= ? and school_id = ?", classId, schoolId).
		Find(&timeTable).Error; err != nil {
		return nil, err
	}

	return timeTable, nil
}

// Update time table by school
func UpdateTimeTableRepository(body *model.TimeTable) (*model.TimeTable, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}

	return body, nil
}

// delete time table by id
func DeleteTimeTableByIdRepository(schoolId string, tableId string) (*model.TimeTable, error) {
	var timeTable *model.TimeTable

	if err := initializer.DB.
		Where("school_id = ? and id = ?", schoolId, tableId).
		First(&timeTable).Error; err != nil {
		return nil, err
	}

	if err := initializer.DB.
		Delete(&timeTable).Error; err != nil {
		return nil, err
	}

	return timeTable, nil
}
