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
	CreatedBy       string                  `json:"createdBy"`
	ModifiedBy      string                  `json:"modifiedBy"`
	ModifiedAt      string                  `json:"modifiedAt"`
}

func NewDoctorObj() *DoctorObj {
	return &DoctorObj{}
}
