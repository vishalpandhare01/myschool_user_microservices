package repository

import (
	"fmt"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
	"gorm.io/gorm"
)

// add class
func AddNewClassRepository(body *model.ClassAndStandrd) (*model.ClassAndStandrd, error) {
	if err := initializer.DB.Create(&body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

// check class already existr
func CheckClassExistRepository(body *model.ClassAndStandrd) bool {
	// Try to find a class or division in the database
	if err := initializer.DB.
		Where("class_name = ? AND division_name = ? AND school_id = ?", body.ClassName, body.DivisionName, body.SchoolID).
		First(&body).Error; err != nil {
		fmt.Println("Error in class: ", err.Error())
		// If no record is found, return false, indicating it doesn't exist
		if err == gorm.ErrRecordNotFound {
			return false // Not found, so class/division doesn't exist
		}
	}
	// Record found, so class/division already exists
	fmt.Println("check")
	return true
}

// get all classes
func GetClassBySchoolIdRepository(schoolId string) (*[]model.ClassAndStandrd, error) {
	var data []model.ClassAndStandrd
	if err := initializer.DB.Where("school_id = ?", schoolId).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// get class by id
func GetClassByIdRepository(id string) (*model.ClassAndStandrd, error) {
	var data model.ClassAndStandrd
	if err := initializer.DB.Where("id = ?", id).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

// delete class by classid and schoolid(userid)
func DeleteClassByIdRepository(classId string, schoolId string) (*model.ClassAndStandrd, error) {
	var data model.ClassAndStandrd

	if err := initializer.DB.Where("id = ? AND school_id = ?", classId, schoolId).Delete(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// check class already exist by classid
func CheckClassExistbyClassIdRepository(classId string, schoolId string) bool {
	var class model.ClassAndStandrd

	if err := initializer.DB.
		Where("id = ? AND school_id =?", classId, schoolId).
		First(&class).Error; err != nil {
		fmt.Println("Error in class: ", err.Error())
		if err == gorm.ErrRecordNotFound {
			return false
		}
	}
	fmt.Println("check")
	return true
}
