package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

type UserBody struct {
	Image             string
	FirstName         string
	LastName          string
	SchoolName        string
	Email             string
	MobileNumber      string
	IsPaidSchool      bool
	Role              string
	Address           string
	SchoolPaymentDate string
}

// create user , school , teacher and student
func CreateUserRepository(body UserBody) (*model.User, error) {
	user := model.User{
		Image:             body.Image,
		FirstName:         body.FirstName,
		LastName:          body.LastName,
		SchoolName:        body.SchoolName,
		Email:             body.Email,
		MobileNumber:      body.MobileNumber,
		IsPaidSchool:      body.IsPaidSchool,
		Role:              body.Role,
		Address:           body.Address,
		SchoolPaymentDate: body.SchoolPaymentDate,
	}

	if err := initializer.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

type OtPBody struct {
	OtpCode int
	IsUsed  bool
	UserId  string
	Attempt int
}

// on send otp create user otp data
func CreateOtpForUserRepository(body OtPBody) (*model.User_Otp, error) {
	var data model.User_Otp

	//if user otp record not exist creating new otp record
	if err := initializer.DB.Where("user_id", body.UserId).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			data.UserId = body.UserId
			data.OtpCode = body.OtpCode
			data.Attempt = 1
			if err := initializer.DB.Create(&data).Error; err != nil {
				return nil, err
			}

			return &data, nil
		}
		return nil, err
	}

	//if user otp record  exist updating existing record
	data.IsUsed = false
	data.UserId = body.UserId
	data.OtpCode = body.OtpCode
	data.Attempt = data.Attempt + 1

	time.AfterFunc(30*time.Second, func() {
		data.IsUsed = true
		if err := initializer.DB.Save(&data).Error; err != nil {
			fmt.Println(err.Error())
		}
	})

	if err := initializer.DB.Save(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil

}

// veryfy otp
func VeryfyOtpRepository(otp int, userId string) (*model.User_Otp, error) {
	var data model.User_Otp

	if err := initializer.DB.Where("user_id = ? AND otp_code = ?", userId, otp).First(&data).Error; err != nil {
		if err.Error() == "record not found" {
			return nil, errors.New("invalid otp")
		}
		return nil, err
	}
	if data.IsUsed {
		return nil, errors.New("otp expired")
	}
	data.IsUsed = true

	if err := initializer.DB.Save(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// get user otp data by user id
func GetOtpDataByUserIdrepository(userId string) (*model.User_Otp, error) {
	var data model.User_Otp

	if err := initializer.DB.Where("user_id", userId).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

// get user by email or phone number
func GetUserByEmailOrPhone(EmailOrPhone string, isEmail bool) (*model.User, error) {
	var data model.User

	if isEmail {
		if err := initializer.DB.Where("email = ?", EmailOrPhone).First(&data).Error; err != nil {
			return nil, err
		}
	} else {
		if err := initializer.DB.Where("mobile_number = ?", EmailOrPhone).First(&data).Error; err != nil {
			return nil, err
		}
	}

	return &data, nil
}

// update user by id
func UpdateUserByIdRepository(body UserBody) (*model.User, error) {
	var user model.User
	if body.Address != "" {
		user.Address = body.Address
	}
	if body.Email != "" {
		user.Email = body.Email
	}
	if body.FirstName != "" {
		user.FirstName = body.FirstName
	}
	if body.LastName != "" {
		user.LastName = body.LastName
	}
	if body.MobileNumber != "" {
		user.MobileNumber = body.MobileNumber
	}
	if body.SchoolName != "" {
		user.SchoolName = body.SchoolName
	}
	if body.Image != "" {
		user.Image = body.Image
	}

	if err := initializer.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
