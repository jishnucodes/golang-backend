package requestData

import (
	// "time"
)

type DoctorObj struct {
	DoctorID        uint      `json:"doctorId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phoneNumber"`
	MobileNumber    string    `json:"mobileNumber"`
	DepartmentName  string    `json:"departmentName"`
	EmployeeID      uint    `json:"employeeId"`
	DepartmentID    uint    `json:"departmentId"`
	Specialty       string    `json:"specialty"`
	ConsultationFee float64   `json:"consultationFee"`
	CreatedAt       string    `json:"createdAt"`
	CreatedBy       string    `json:"createdBy"`
	ModifiedBy      string    `json:"modifiedBy"`
	ModifiedAt      string    `json:"modifiedAt"`
}

func NewDoctorObj() *DoctorObj {
	return &DoctorObj{}
}
