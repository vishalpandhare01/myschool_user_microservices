package repository

import (
	"fmt"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/funcation"
)

// add student in school
func AddNewStudentRepository(body *model.Student) (*model.Student, error) {
	var user *model.User
	if err := initializer.DB.Create(&body).Error; err != nil {
		userDelete := initializer.DB.Where("id = ?", body.UserID).Delete(&user)
		if userDelete != nil {
			fmt.Println("User not register properly hence delete: ", err.Error())
		}
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

// get  student by id
func GetStudentRepository(studentId string) (*model.Student, error) {
	var data *model.Student
	if err := initializer.DB.
		Where("id = ?", studentId).
		Preload("User").
		Preload("StudentClass").
		Preload("School").
		First(&data).
		Error; err != nil {
		return nil, err
	}

	return data, nil
}
