package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string    `gorm:"type:char(36);primarykey"`
	Image        string    `gorm:"type:text;"`
	FirstName    string    `gorm:"type:varchar(255);"`
	LastName     string    `gorm:"type:varchar(255);"`
	SchoolName   string    `gorm:"type:varchar(255);"`
	Email        string    `gorm:"type:varchar(255);not null;unique"`
	MobileNumber string    `gorm:"type:varchar(20);not null;unique"`
	IsPaidSchool bool      `gorm:"type:boolean;default:false"`
	Role         string    `gorm:"type:enum('admin', 'teacher', 'school', 'student');not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *User) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}
