package models

import (
	"time"
)

// Department represents a department in the clinic
type CMSDepartments struct {
	DepartmentID     uint      `gorm:"column:DepartmentID;primaryKey;autoIncrement"`
	DepartmentName   string    `gorm:"column:DepartmentName;type:varchar(100);not null"`
	HeadOfDepartment uint      `gorm:"column:HeadOfDepartment"`
	ContactNumber    string    `gorm:"column:ContactNumber;type:varchar(20)"`
	Email            string    `gorm:"column:Email;type:varchar(100)"`
	Location         string    `gorm:"column:Location;type:varchar(100)"`
	OperatingHours   string    `gorm:"column:OperatingHours;type:varchar(100)"`
	NumberOfStaff    int       `gorm:"column:NumberOfStaff;default:0"`
	ServicesOffered  string    `gorm:"column:ServicesOffered;type:nvarchar(max)"`
	Status           int       `gorm:"column:Status"`
	CreatedAt        time.Time `gorm:"column:CreatedAt"`
	CreatedBy        uint    `gorm:"column:CreatedBy;type:int"`
	ModifiedAt       time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy       uint    `gorm:"column:ModifiedBy;type:int"`

	EmployeeMaster   *CMSEmployeeMaster `gorm:"foreignKey:HeadOfDepartment"`
}

// TableName specifies the table name for the Department model
func (CMSDepartments) TableName() string {
	return "CMS_Departments"
}
