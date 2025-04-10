package models

import (
	"time"
)

// CMSDoctor represents the CMS_Doctors table
type CMSDoctor struct {
	DoctorID        uint      `gorm:"primaryKey;autoIncrement;column:DoctorID"`
	FirstName       string    `gorm:"column:FirstName"`
	LastName        string    `gorm:"column:LastName"`
	Specialty       string    `gorm:"column:Specialty"`
	ContactNumber   string    `gorm:"column:ContactNumber"`
	Email           string    `gorm:"column:Email;unique"`
	ConsultationFee float64   `gorm:"column:ConsultationFee;type:decimal(10,2)"`
	CreatedAt       time.Time `gorm:"column:CreatedAt"`
	CreatedBy       string    `gorm:"column:CreatedBy"`
	ModifiedBy      string    `gorm:"column:ModifiedBy"`
	ModifiedAt      time.Time `gorm:"column:ModifiedAt"`
}

// TableName overrides the default table name
func (CMSDoctor) TableName() string {
	return "CMS_Doctors"
}
