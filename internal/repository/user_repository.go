package repository

import (
	"github.com/vishalpandhare01/initializer"
	"github.com/vishalpandhare01/internal/model"
)

type UserBody struct {
	Image        string
	FirstName    string
	LastName     string
	SchoolName   string
	Email        string
	MobileNumber string
	IsPaidSchool string
	Role         string
}

func CreateUserRepository(body UserBody) (*model.User, error) {
	var user model.User
	if err := initializer.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
