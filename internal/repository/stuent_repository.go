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

	fmt.Println("schoolId...........", schoolId)
	fmt.Println("mobileNumber...........", mobileNumber)
	fmt.Println("registerNumber...........", registerNumber)
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
	}
	if registerNumber != "" {
		query = query.Where("register_number = ? AND school_id = ?", registerNumber, schoolId)
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

// get  student by id
func GetStudentRepository(userId string) (*model.Student, error) {
	var data *model.Student
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

// update student by school
func UpdateStudentRepository(body *model.Student) (*model.Student, error) {
	var student *model.Student
	var user *model.User
	if err := initializer.DB.Where("id = ?", body.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	if err := initializer.DB.
		Where("user_id = ?", body.UserID).
		First(&student).
		Error; err != nil {
		return nil, err
	}

	if body.Address != "" {
		student.Address = body.Address
		user.Address = body.Address
	}
	if body.ClassID != "" {
		student.ClassID = body.ClassID
	}
	if body.ConductInSchool != "" {
		student.ConductInSchool = body.ConductInSchool
	}
	if body.DateOfAddmission != "" {
		student.DateOfAddmission = body.DateOfAddmission
	}
	if body.DateOfBirth != "" {
		student.DateOfBirth = body.DateOfBirth
	}
	if body.DateOfLeaving != "" {
		student.DateOfLeaving = body.DateOfLeaving
	}
	if body.Email != "" {
		student.Email = body.Email
		user.Email = body.Email
	}
	if body.FatherName != "" {
		student.FatherName = body.FatherName
	}
	if body.FirstName != "" {
		student.FirstName = body.FirstName
	}
	if body.LastName != "" {
		student.LastName = body.LastName
	}
	if body.Image != "" {
		student.Image = body.Image
	}
	if body.MobileNumber != "" {
		student.MobileNumber = body.MobileNumber
		user.MobileNumber = body.MobileNumber
	}
	if body.MotherName != "" {
		student.MotherName = body.MotherName
	}

	if body.ParentsMobileNumber != "" {
		student.ParentsMobileNumber = body.ParentsMobileNumber
	}
	if body.PlaceOfBirth != "" {
		student.PlaceOfBirth = body.PlaceOfBirth
	}

	if body.ProgrssInStudies != "" {
		student.ProgrssInStudies = body.ProgrssInStudies
	}
	if body.RegisterNumber != 0 {
		student.RegisterNumber = body.RegisterNumber
	}
	if body.ReasonOfLeaving != "" {
		student.ReasonOfLeaving = body.ReasonOfLeaving
	}
	if body.ReligionOrCast != "" {
		student.ReligionOrCast = body.ReligionOrCast
	}
	if body.Remark != "" {
		student.Remark = body.Remark
	}
	if body.Gender != "" {
		student.Gender = body.Gender
	}
	if body.AcademicYear != "" {
		student.AcademicYear = body.AcademicYear
	}

	if err := initializer.DB.
		Save(&user).
		Error; err != nil {
		return nil, err
	}
	if err := initializer.DB.
		Save(&student).
		Error; err != nil {
		return nil, err
	}
	return student, nil
}

// remove student from school and save stdent data
func RemoveStudentFromSchoolRepository(userId string) (*model.PastStudent, error) {
	var data *model.Student
	var user *model.User
	if err := initializer.DB.
		Where("user_id = ?", userId).
		First(&data).
		Error; err != nil {
		return nil, err
	}
	if err := initializer.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}

	var pastStudent = model.PastStudent{
		UserID:              data.UserID,
		ClassID:             data.ClassID,
		SchoolID:            data.SchoolID,
		Image:               data.Image,
		FirstName:           data.FirstName,
		LastName:            data.LastName,
		SchoolName:          data.SchoolName,
		Address:             data.Address,
		Email:               data.Email,
		MobileNumber:        data.MobileNumber,
		IsPaidSchool:        data.IsPaidSchool,
		SchoolPaymentDate:   data.SchoolPaymentDate,
		Role:                data.Role,
		RegisterNumber:      data.RegisterNumber,
		MotherName:          data.MotherName,
		FatherName:          data.FatherName,
		DateOfBirth:         data.DateOfBirth,
		PlaceOfBirth:        data.PlaceOfBirth,
		DateOfAddmission:    data.DateOfAddmission,
		DateOfLeaving:       data.DateOfLeaving,
		ReligionOrCast:      data.ReligionOrCast,
		ProgrssInStudies:    data.ProgrssInStudies,
		ConductInSchool:     data.ConductInSchool,
		Remark:              data.Remark,
		IsLeaved:            data.IsLeaved,
		StudentDoc:          data.StudentDoc,
		ReasonOfLeaving:     data.ReasonOfLeaving,
		ParentsMobileNumber: data.ParentsMobileNumber,
	}

	fmt.Println("Here..........")
	if err := initializer.DB.Create(&pastStudent).Error; err != nil {
		return nil, err
	}

	if err := initializer.DB.Delete(&data).Error; err != nil {
		return nil, err
	}
	if err := initializer.DB.Delete(&user).Error; err != nil {
		return nil, err
	}

	return &pastStudent, nil
}

// moves all students from one class to another (batch update).
func MoveBulkStudentToAnotherClassRepository(currentClassId string, nextClassId string, schoolId string) (*[]model.Student, error) {
	var students []model.Student

	// Update the students' class_id in bulk
	if err := initializer.DB.Model(&model.Student{}).
		Where("class_id = ? AND school_id = ?", currentClassId, schoolId).
		Update("class_id", nextClassId).Error; err != nil {
		return nil, err
	}

	// Fetch the updated list of students to return
	if err := initializer.DB.Where("class_id = ?", nextClassId).Find(&students).Error; err != nil {
		return nil, err
	}

	return &students, nil
}
