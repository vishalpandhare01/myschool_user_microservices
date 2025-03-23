package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeTable struct {
	ID              string           `gorm:"type:char(36);primarykey;not null"`
	ClassID         string           `gorm:"type:varchar(36);not:null"`
	TeacherID       string           `gorm:"type:varchar(36);not:null"` //staff id
	SchoolID        string           `gorm:"type:varchar(36);not:null"`
	DayOfWeek       string           `gorm:"type:varchar(10);not:null"`
	Subject         string           `gorm:"type:varchar(100);not:null"`
	StartTime       string           `gorm:"type:varchar(10);not:null"`
	EdnTime         string           `gorm:"type:varchar(10);not:null"`
	IsExamTimeTable bool             `gorm:"type:boolean;default false"`
	CreatedAt       time.Time        `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt       time.Time        `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	School          *User            `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	Class           *ClassAndStandrd `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;"`
	Teacher         *Staff           `gorm:"foreignKey:TeacherID;constraint:OnDelete:CASCADE;"`
}

func (U *TimeTable) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}
