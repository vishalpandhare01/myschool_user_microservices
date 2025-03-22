package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FeeType struct {
	ID        string    `gorm:"type:char(36);primarykey"`
	FeeName   string    `gorm:"type:varchar(255);"`
	SchoolID  string    `gorm:"type:varchar(36);not:null"`
	School    *User     `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *FeeType) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

type FeesStructure struct {
	ID           string    `gorm:"type:char(36);primarykey"`
	SchoolID     string    `gorm:"type:varchar(36);not:null"`
	FeeTypeID    string    `gorm:"type:varchar(36);not:null"`
	Amount       float64   `gorm:"type:decimal(10,2);not:null"`
	School       *User     `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	FeeType      *FeeType  `gorm:"foreignKey:FeeTypeID;constraint:OnDelete:CASCADE;"`
	AcademicYear string    `gorm:"type:varchar(20);not null" json:"academicYear"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

func (U *FeesStructure) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

type StudentFees struct {
	ID              string    `gorm:"type:char(36);primarykey"`
	UserID          string    `gorm:"type:varchar(36);not:null"`
	SchoolID        string    `gorm:"type:varchar(36);not:null"`
	AcademicYear    string    `gorm:"type:varchar(10);not:null"` // Store "YYYY-YYYY" format
	Status          string    `gorm:"type:enum('paid','pending','partial');default:'pending'"`
	TotalAmount     float64   `gorm:"type:decimal(10,2);not:null"`
	PaidAmount      float64   `gorm:"type:decimal(10,2);not:null"`
	RemainingAmount float64   `gorm:"type:decimal(10,2)" json:"remainingAmount,omitempty"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	User            *User     `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	FeesDetails     string    `gorm:"text;not:null"`
}

func (U *StudentFees) BeforeCreate(tx *gorm.DB) (err error) {
	U.ID = uuid.New().String()
	return
}

type Payment struct {
	UserID    string    `gorm:"type:varchar(36);primarykey;not:null"`
	SchoolID  string    `gorm:"type:varchar(36);not:null"`
	Amount    float64   `gorm:"type:decimal(10,2);not:null"`
	Reference string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}
