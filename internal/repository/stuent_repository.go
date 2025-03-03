package repository

import (
	"fmt"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/funcation"
)

// add student
func AddNewStudentRepository(body *model.Student) (*model.Student, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

// get all students
func GetAllStudentRepository(pageStr string, limitStr string, schoolId string) (interface{}, error) {
	var data *[]model.Student
	var totalData []model.Student
	offset, limitInt := funcation.Pagination(pageStr, limitStr)
	if err := initializer.DB.
		Where("school_id = ?", schoolId).
		Preload("User").
		Preload("StudentClass").
		Preload("School").
		Limit(limitInt).
		Offset(offset).
		Find(&data).
		Order("id DESC").
		Error; err != nil {
		return nil, err
	}
	if err := initializer.DB.
		Where("school_id = ?", schoolId).
		Find(&totalData).
		Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	responseData := utils.SuccessListResponse{
		Total:   len(totalData),
		Perpage: limitInt,
		Page:    offset,
		Data:    data,
	}
	return responseData, nil
}
