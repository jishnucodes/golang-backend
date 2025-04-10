package requestData

import (
	"time"
)

type DoctorObj struct {
	DoctorID        uint      `json:"doctorId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Specialty       string    `json:"specialty"`
	ContactNumber   string    `json:"contactNumber"`
	Email           string    `json:"email"`
	ConsultationFee float64   `json:"consultationFee"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       string    `json:"createdBy"`
	ModifiedBy      string    `json:"modifiedBy"`
	ModifiedAt      time.Time `json:"modifiedAt"`
}

func NewDoctorObj() *DoctorObj {
	return &DoctorObj{}
}
