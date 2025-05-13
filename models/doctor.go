package models

import (
	"time"
)

// CMSDoctor represents the CMS_Doctors table
type CMSDoctor struct {
	DoctorID        uint      `gorm:"primaryKey;autoIncrement;column:DoctorID"`
	EmployeeID      uint    `gorm:"column:EmployeeID"`
	Specialty       string    `gorm:"column:Specialty"`
	ConsultationFee float64   `gorm:"column:ConsultationFee;type:decimal(10,2)"`
	CreatedAt       time.Time `gorm:"column:CreatedAt"`
	CreatedBy       uint    `gorm:"column:CreatedBy"`
	ModifiedBy      uint    `gorm:"column:ModifiedBy"`
	ModifiedAt      time.Time `gorm:"column:ModifiedAt"`

	EmployeeMaster *CMSEmployeeMaster `gorm:"foreignKey:EmployeeID"`
}

// TableName overrides the default table name
func (CMSDoctor) TableName() string {
	return "CMS_Doctors"
}
