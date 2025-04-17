package requestData

import "time"

type EmployeeObj struct {
	EmployeeID    uint      `json:"employeeId"`
	EmployeeCode  string    `json:"employeeCode"`
	UserID        uint      `json:"userId"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"` 
	Email         string    `json:"email"`
	PhoneNumber   string    `json:"phoneNumber"`
	MobileNumber  string    `json:"mobileNumber"`
	Address       string    `json:"address"`
	BloodGroup    string    `json:"bloodGroup"`
	HireDate      time.Time `json:"hireDate"`
	JobTitle      string    `json:"jobTitle"`
	DepartmentID  uint      `json:"departmentId"`
	EmployeeType  string    `json:"employeeType"`
	CreatedAt     time.Time `json:"createdAt"`
	CreatedBy     string    `json:"createdBy"`
	ModifiedAt    time.Time `json:"modifiedAt"`
	ModifiedBy    string    `json:"modifiedBy"`
}

func NewEmployeeObj() *EmployeeObj {
	return &EmployeeObj{}
}

