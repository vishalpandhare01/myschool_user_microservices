package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID                string    `gorm:"type:char(36);primarykey"`
	Image             string    `gorm:"type:text;"`
	FirstName         string    `gorm:"type:varchar(255);"`
	LastName          string    `gorm:"type:varchar(255);"`
	SchoolName        string    `gorm:"type:varchar(255);"`
	Address           string    `gormL:"type:text;"`
	Email             string    `gorm:"type:varchar(255);not null;unique"`
	MobileNumber      string    `gorm:"type:varchar(20);not null;unique"`
	IsPaidSchool      bool      `gorm:"type:boolean;default:false"`
	SchoolPaymentDate string    `gormL:"type:text;"`
	Role              string    `gorm:"type:enum('admin', 'staff', 'school', 'student');not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *User) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

type User_Otp struct {
	ID        string    `gorm:"primaryKey;type:char(36);not null"`
	OtpCode   int       `gorm:"type:char(6);not:null"`
	IsUsed    bool      `gorm:"bool;default:false"`
	UserId    string    `gorm:"type:varchar(36);not:null"`
	Attempt   int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *User_Otp) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}
