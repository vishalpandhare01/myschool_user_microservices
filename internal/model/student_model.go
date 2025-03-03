package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID                  string           `gorm:"type:char(36);primarykey"`
	UserID              string           `gorm:"type:varchar(36);not:null"`
	ClassID             string           `gorm:"type:varchar(36);not:null"`
	SchoolID            string           `gorm:"type:varchar(36);not:null"` //user_id
	RegisterNumber      int64            `gorm:"type:bigint;"`
	MotherName          string           `gorm:"type:varchar(255);"`
	FatherName          string           `gorm:"type:varchar(255);"`
	DateOfBirth         string           `gorm:"type:varchar(255);"`
	PlaceOfBirth        string           `gorm:"type:varchar(255);"`
	DateOfAddmission    string           `gorm:"type:varchar(255);"`
	DateOfLeaving       string           `gorm:"type:varchar(255);"`
	ReligionOrCast      string           `gorm:"type:varchar(255);"`
	ProgrssInStudies    string           `gorm:"type:varchar(255);"`
	ConductInSchool     string           `gorm:"type:varchar(255);"`
	Remark              string           `gorm:"type:varchar(255);"`
	IsLeaved            bool             `gorm:"type:boolean;default:false"`
	StudentDoc          string           `gorm:"type:text;"`
	ReasonOfLeaving     string           `gorm:"type:text;"`
	ParentsMobileNumber string           `gorm:"type:varchar(20);not null;unique"`
	CreatedAt           time.Time        `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	User                *User            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	School              *User            `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	StudentClass        *ClassAndStandrd `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;"`
}

func (U *Student) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}
