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
func GetAllStudentRepository(
	pageStr string,
	limitStr string,
	schoolId string,
	mobileNumber string,
	registerNumber string,
	email string,
	classID string,
	fName string,
	lName string,
) (interface{}, error) {
	var data *[]model.Student
	var totalData []model.Student
	offset, limitInt := funcation.Pagination(pageStr, limitStr)

	query := initializer.DB.Where("school_id = ?", schoolId)
	if email != "" {
		query = query.Where("email = ?", email)
	}
	if classID != "" {
		query = query.Where("class_id = ?", classID)
	}
	if fName != "" {
		query = query.Where("first_name = ?", fName)
	}
	if fName != "" {
		query = query.Where("last_name = ?", lName)
	}
	if mobileNumber != "" {
		query = query.Where("mobile_number = ?", mobileNumber)
	}
	if registerNumber != "" {
		query = query.Where("register_number = ?", registerNumber)
	}

	query = query.Preload("User").
		Preload("StudentClass").
		Preload("School").
		Limit(limitInt).
		Offset(offset).
		Find(&data).
		Order("id DESC")

	if err := query.Error; err != nil {
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

// update student by school
func UpdateStudentRepository(body *model.Student) (*model.Student, error) {
	if err := initializer.DB.Save(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}
