package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Attendance struct {
	ID              string           `gorm:"type:char(36);primarykey;not:null"`
	ClassID         string           `gorm:"type:char(36);not:null"`
	StudentID       string           `gorm:"type:char(36);not:null"`
	TeacherID       string           `gorm:"type:char(36);not:null"`
	SchoolID        string           `gorm:"type:char(36);not:null"`
	Subject         string           `gorm:"type:varchar(100);not:null"`
	Date            string           `gorm:"type:varchar(10);not:null"`
	Status          string           `gorm:"type:enum('present', 'absent', 'leave');not:null"`
	School          *User            `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	ClassAndStandrd *ClassAndStandrd `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;"`
	Student         *Student         `gorm:"foreignKey:StudentID;constraint:OnDelete:CASCADE;"`
	Teacher         *Staff           `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE;"`
	CreatedAt       time.Time        `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt       time.Time        `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (A *Attendance) BeforeCreate(tx *gorm.DB) (err error) {
	A.ID = uuid.New().String()
	return
}
