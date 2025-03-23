package repository

import (
	"fmt"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/funcation"
)

// add staff in school
func AddNewStaffRepository(body *model.Staff) (*model.Staff, error) {
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

// get all staff
func GetAllStaffRepository(
	pageStr string,
	limitStr string,
	schoolId string,
	mobileNumber string,
	email string,
	classID string,
	fName string,
	lName string,
) (interface{}, error) {
	var data *[]model.Staff
	var totalData []model.Staff
	offset, limitInt := funcation.Pagination(pageStr, limitStr)

	fmt.Println("schoolId...........", schoolId)
	fmt.Println("mobileNumber...........", mobileNumber)
	fmt.Println("email...........", email)
	fmt.Println("classID...........", classID)
	fmt.Println("fName...........", fName)
	fmt.Println("lName...........", lName)

	query := initializer.DB
	if email != "" {
		query = query.Where("email = ? AND school_id = ?", email, schoolId)
	}
	if classID != "" {
		query = query.Where("class_id = ? AND school_id = ?", classID, schoolId)
	}
	if fName != "" {
		query = query.Where("first_name = ? AND school_id = ?", fName, schoolId)
	}
	if lName != "" {
		query = query.Where("last_name = ? AND school_id = ?", lName, schoolId)
	}
	if mobileNumber != "" {
		query = query.Where("mobile_number = ? AND school_id = ?", mobileNumber, schoolId)
	} else {
		query = query.Where("school_id = ?", schoolId)
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

	fmt.Println(query)

	responseData := utils.SuccessListResponse{
		Total:   len(totalData),
		Perpage: limitInt,
		Page:    offset,
		Data:    data,
	}
	return responseData, nil
}

// get staff by userId
func GetStaffRepository(userId string) (*model.Staff, error) {
	var data *model.Staff
	if err := initializer.DB.
		Where("user_id = ?", userId).
		Preload("User").
		Preload("StudentClass").
		Preload("School").
		First(&data).
		Error; err != nil {
		return nil, err
	}

	return data, nil
}

// check staff exist with schoolid staff id
func CheckStaffExistStaffRepository(userId string, schoolID string) bool {
	var data *model.Staff
	if err := initializer.DB.
		Where("user_id = ? AND school_id = ?", userId, schoolID).
		First(&data).
		Error; err != nil {
		return false
	}

	return true
}

// update staff by userId
func UpdateStaffRepository(body *model.Staff) (*model.Staff, error) {
	var Staff *model.Staff
	var user *model.User
	if err := initializer.DB.Where("id = ?", body.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	if err := initializer.DB.
		Where("user_id = ?", body.UserID).
		First(&Staff).
		Error; err != nil {
		return nil, err
	}

	if body.Address != "" {
		Staff.Address = body.Address
		user.Address = body.Address
	}
	if body.ClassID != "" {
		Staff.ClassID = body.ClassID
	}

	if body.Email != "" {
		Staff.Email = body.Email
		user.Email = body.Email
	}

	if body.FirstName != "" {
		Staff.FirstName = body.FirstName
	}
	if body.LastName != "" {
		Staff.LastName = body.LastName
	}
	if body.Image != "" {
		Staff.Image = body.Image
	}
	if body.MobileNumber != "" {
		Staff.MobileNumber = body.MobileNumber
		user.MobileNumber = body.MobileNumber
	}

	if body.Gender != "" {
		Staff.Gender = body.Gender
	}

	if err := initializer.DB.
		Save(&user).
		Error; err != nil {
		return nil, err
	}
	if err := initializer.DB.
		Save(&Staff).
		Error; err != nil {
		return nil, err
	}
	return Staff, nil
}

// delete staff account
func DeleteStaffAcountSchoolRepository(userId string) (*model.User, error) {
	var data *model.Staff
	var user *model.User
	if err := initializer.DB.
		Where("user_id = ?", userId).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	if err := initializer.
		DB.Where("id = ?", userId).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	if err := initializer.DB.Delete(&data).Error; err != nil {
		return nil, err
	}
	if err := initializer.DB.Delete(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
