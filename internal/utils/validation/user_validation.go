package validation

import (
	"fmt"
	"regexp"
	"time"

	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

var user model.User

func CheckNameExist(name string) bool {
	if err := initializer.DB.Where("name = ?", name).First(&user).Error; err != nil {
		fmt.Println("Name Not Exist: ", err.Error())
		return false
	}
	return true
}

func CheckEmailExist(email string) bool {
	if err := initializer.DB.Where("email = ?", email).First(&user).Error; err != nil {
		fmt.Println("Email Not Exist: ", err.Error())
		return false
	}
	return true
}

func CheckMobileNumberExist(mobileNumber string) bool {
	if err := initializer.DB.Where("mobile_number = ?", mobileNumber).First(&user).Error; err != nil {
		fmt.Println("Mobile Number Not Exist: ", err.Error())
		return false
	}
	return true
}

func CheckRoleIsCorrect(Role string) bool {
	if Role == "admin" {
		if err := initializer.DB.Where("role = ?", Role).First(&user).Error; err != nil {
			fmt.Println("Admin Not Exist: ", err.Error())
			return true
		}
		return false
	} else if Role == "teacher" {
		return true
	} else if Role == "school" {
		return true
	} else if Role == "student" {
		return true
	}
	return false
}

func ValidDateOfBirth(date string) bool {
	layout := "2006-01-02"
	_, err := time.Parse(layout, date)
	return err == nil
}

func ValidateYearRange(input string) bool {
	// Regular expression to match the format YYYY-YYYY
	// ^: start of the string
	// \d{4}: exactly four digits
	// -: a literal hyphen
	// \d{4}: exactly four digits
	// $: end of the string
	re := regexp.MustCompile(`^\d{4}-\d{4}$`)

	return re.MatchString(input)
}
