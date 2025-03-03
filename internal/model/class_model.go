package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassAndStandrd struct {
	ID           string    `gorm:"type:char(36);primarykey"`
	ClassName    string    `gorm:"type:varchar(255);not:null"`
	DivisionName string    `gorm:"type:varchar(255);not:null"`
	SchoolID     string    `gorm:"type:varchar(36);not:null"` //user_id
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	School       *User     `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
}

func (U *ClassAndStandrd) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}
