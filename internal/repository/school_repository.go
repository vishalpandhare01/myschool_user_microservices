package repository

import (
	"fmt"
	"time"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
	"github.com/vishalpandhare01/internal/utils"
	"github.com/vishalpandhare01/internal/utils/funcation"
)

//getschools // all school ,search by name , pagination , paid or not filter ,

func GetSchoolsRepository(pageStr string, limitStr string, school_name string, isPaid string) (interface{}, error) {
	var data []model.User
	var totalData []model.User

	offset, limitInt := funcation.Pagination(pageStr, limitStr)
	fmt.Println("limitInt", limitInt, "offset", offset)

	query := initializer.DB

	if school_name != "" && isPaid != "" {
		if isPaid == "true" {
			query = query.Where("role = ? AND school_name LIKE ? AND is_paid_school = ?", "school", "%"+school_name+"%", true)
		} else {
			query = query.Where("role = ? AND school_name LIKE ? AND is_paid_school = ?", "school", "%"+school_name+"%", false)
		}
	} else if school_name != "" {
		query = query.Where("role = ? AND school_name LIKE ?", "school", "%"+school_name+"%")
	} else if isPaid != "" {
		if isPaid == "true" {
			query = query.Where("role = ? AND school_name LIKE ? AND is_paid_school = ?", "school", "%"+school_name+"%", true)
		} else {
			query = query.Where("role = ? AND school_name LIKE ? AND is_paid_school = ?", "school", "%"+school_name+"%", false)
		}
	} else {
		query = query.Where("role = ?", "school")
	}

	if err := query.Find(&totalData).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// offset = (offset - 1) * limitInt
	query.Limit(limitInt).Offset(offset).Order("id DESC")
	if err := query.Find(&data).Error; err != nil {
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

func UpdateSchoolPaymentRepository(schoolId string) (*model.User, error) {
	var user model.User
	if err := initializer.DB.Where("id = ? and role = ?", schoolId, "school").First(&user).Error; err != nil {
		return nil, err
	}

	user.IsPaidSchool = !user.IsPaidSchool
	currentTime := time.Now()
	currentDate := currentTime.Format("2006-01-02")
	user.SchoolPaymentDate = currentDate

	if err := initializer.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetSchoolByIdRepository(userId string) (*model.User, error) {
	var user model.User
	if err := initializer.DB.Where("id = ? and role = ?", userId, "school").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
