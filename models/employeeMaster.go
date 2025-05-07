package models

import "time"

// CMSEmployee represents the CMS_Employees table
type CMSEmployeeMaster struct {
	EmployeeID    uint      `gorm:"primaryKey;autoIncrement;column:EmployeeID"`
	EmployeeCode  string    `gorm:"column:EmployeeCode;type:varchar(50)"`
	UserID        uint      `gorm:"column:UserID"`
	FirstName     string    `gorm:"column:FirstName;type:varchar(50)"`
	LastName      string    `gorm:"column:LastName;type:varchar(50)"`
	Email         string    `gorm:"column:Email;type:varchar(100);unique"`
	PhoneNumber   string    `gorm:"column:PhoneNumber;type:varchar(20)"`
	MobileNumber  string    `gorm:"column:MobileNumber;type:varchar(20)"`
	Address       string    `gorm:"column:Address;type:nvarchar(max)"`
	BloodGroup    string    `gorm:"column:BloodGroup;type:varchar(10)"`
	HireDate      time.Time `gorm:"column:HireDate"`
	JobTitle      string    `gorm:"column:JobTitle;type:varchar(50)"`
	DepartmentID  uint      `gorm:"column:DepartmentID"`
	EmployeeType  uint      `gorm:"column:EmployeeType"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
	CreatedBy     string    `gorm:"column:CreatedBy;type:varchar(50)"`
	ModifiedAt    time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy    string    `gorm:"column:ModifiedBy;type:varchar(50)"`

	User *CMSUser `gorm:"foreignKey:UserID"`
	Department *CMSDepartments `gorm:"foreignKey:DepartmentID"`
}

// TableName specifies the table name for the Employee model
func (CMSEmployeeMaster) TableName() string {
	return "CMS_EmployeeMaster"
}
