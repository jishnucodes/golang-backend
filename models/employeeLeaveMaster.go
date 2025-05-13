package models

import "time"

// CMSEmployee represents the CMS_Employees table
type CMSEmployeeLeaveMaster struct {
	ID         uint      `gorm:"primaryKey;autoIncrement;column:ID"`
	EmployeeID uint      `gorm:"column:EmployeeID"`
	DateFrom   time.Time `gorm:"column:DateFrom"`
	DateTo     time.Time `gorm:"column:DateTo"`
	TimeFrom   time.Time `gorm:"column:TimeFrom"`
	TimeTo     time.Time `gorm:"column:TimeTo"`
	Approved   uint      `gorm:"column:Approved"`
	Remarks    string    `gorm:"column:Remarks"`
	LeaveType  uint      `gorm:"column:LeaveType"`
	CreatedAt  time.Time `gorm:"column:CreatedAt"`
	CreatedBy  uint    `gorm:"column:CreatedBy"`
	ModifiedAt time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy uint    `gorm:"column:ModifiedBy"`

	EmployeeMaster *CMSEmployeeMaster `gorm:"foreignKey:EmployeeID"`
}

// TableName specifies the table name for the Employee model
func (CMSEmployeeLeaveMaster) TableName() string {
	return "CMS_EmployeeLeaveMaster"
}
