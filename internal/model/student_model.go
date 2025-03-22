package model

import (
	"time"
)

type Student struct {
	UserID              string           `gorm:"type:varchar(36);primarykey;not:null"`
	ClassID             string           `gorm:"type:varchar(36);not:null"`
	SchoolID            string           `gorm:"type:varchar(36);not:null"` //user_id
	Image               string           `gorm:"type:text;"`
	FirstName           string           `gorm:"type:varchar(255);"`
	LastName            string           `gorm:"type:varchar(255);"`
	SchoolName          string           `gorm:"type:varchar(255);"`
	Address             string           `gormL:"type:text;"`
	Email               string           `gorm:"type:varchar(255);"`
	MobileNumber        string           `gorm:"type:varchar(20);"`
	IsPaidSchool        bool             `gorm:"type:boolean;default:false"`
	SchoolPaymentDate   string           `gormL:"type:text;"`
	Role                string           `gorm:"type:enum('student');default:'student'"`
	Gender              string           `gorm:"type:enum('male','female','other');default:'other'"`
	AcademicYear        string           `gorm:"type:varchar(10);"` // Store "YYYY-YYYY" format
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
	ParentsMobileNumber string           `gorm:"type:varchar(20)"`
	CreatedAt           time.Time        `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt           time.Time        `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
	User                *User            `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	School              *User            `gorm:"foreignKey:SchoolID;constraint:OnDelete:CASCADE;"`
	StudentClass        *ClassAndStandrd `gorm:"foreignKey:ClassID;constraint:OnDelete:CASCADE;"`
}

type PastStudent struct {
	ID                  string    `gorm:"type:char(36);"`
	UserID              string    `gorm:"type:varchar(36);not:null"`
	ClassID             string    `gorm:"type:varchar(36);not:null"`
	SchoolID            string    `gorm:"type:varchar(36);not:null"` //user_id
	Image               string    `gorm:"type:text;"`
	FirstName           string    `gorm:"type:varchar(255);"`
	LastName            string    `gorm:"type:varchar(255);"`
	SchoolName          string    `gorm:"type:varchar(255);"`
	Address             string    `gormL:"type:text;"`
	Email               string    `gorm:"type:varchar(255);"`
	MobileNumber        string    `gorm:"type:varchar(20);"`
	IsPaidSchool        bool      `gorm:"type:boolean;default:false"`
	SchoolPaymentDate   string    `gormL:"type:text;"`
	Role                string    `gorm:"type:enum('student');default:'student'"`
	RegisterNumber      int64     `gorm:"type:bigint;"`
	MotherName          string    `gorm:"type:varchar(255);"`
	FatherName          string    `gorm:"type:varchar(255);"`
	DateOfBirth         string    `gorm:"type:varchar(255);"`
	PlaceOfBirth        string    `gorm:"type:varchar(255);"`
	DateOfAddmission    string    `gorm:"type:varchar(255);"`
	DateOfLeaving       string    `gorm:"type:varchar(255);"`
	ReligionOrCast      string    `gorm:"type:varchar(255);"`
	ProgrssInStudies    string    `gorm:"type:varchar(255);"`
	ConductInSchool     string    `gorm:"type:varchar(255);"`
	Remark              string    `gorm:"type:varchar(255);"`
	IsLeaved            bool      `gorm:"type:boolean;default:false"`
	StudentDoc          string    `gorm:"type:text;"`
	ReasonOfLeaving     string    `gorm:"type:text;"`
	ParentsMobileNumber string    `gorm:"type:varchar(20);"`
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}
