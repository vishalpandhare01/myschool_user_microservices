package model

import (
	"time"
)

type Staff struct {
	UserID         string           `gorm:"type:varchar(36);primarykey;not:null"`
	ClassID        string           `gorm:"type:varchar(36);not:null"`
	SchoolID       string           `gorm:"type:varchar(36);not:null"`
	Image          string           `gorm:"type:text;"`
	FirstName      string           `gorm:"type:varchar(255);"`
	LastName       string           `gorm:"type:varchar(255);"`
	SpecialSubject string           `gorm:"type:varchar(255);"`
	Gender         string           `gorm:"type:enum('male','female','other');default:'other'"`
	Address        string           `gormL:"type:text;"`
	Email          string           `gorm:"type:varchar(255);"`
	MobileNumber   string           `gorm:"type:varchar(20);"`
	Role           string           `gorm:"type:enum('teacher','accountant','staff','hod');default:'staff'"`
	CreatedAt      time.Time        `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt      time.Time        `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	User           *User            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	School         *User            `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	StudentClass   *ClassAndStandrd `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;"`
}
