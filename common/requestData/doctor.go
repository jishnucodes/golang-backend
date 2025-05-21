package requestData

import (
// "time"
)

type DoctorObj struct {
	DoctorID        uint                    `json:"doctorId"`
	EmployeeID      uint                    `json:"employeeId"`
	Specialty       string                  `json:"specialty"`
	ConsultationFee float64                 `json:"consultationFee"`
	Employee        EmployeeObj             `json:"employee"`
	Availabilities  []DoctorAvailabilityObj `json:"availabilities"`
	CreatedAt       string                  `json:"createdAt"`
	CreatedBy       uint                    `json:"createdBy"`
	ModifiedBy      uint                    `json:"modifiedBy"`
	ModifiedAt      string                  `json:"modifiedAt"`
}

type ListDoctorByMonthSearchQuery struct {
	DepartmentId *int `form:"departmentId" json:"departmentId"`
	Month        *int `form:"month" json:"month"`
	Year         *int `form:"year" json:"year"`
}

func NewDoctorObj() *DoctorObj {
	return &DoctorObj{}
}
