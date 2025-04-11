package requestData

import (
	"time"
)

type DoctorAvailabilityObj struct {
	AvailabilityID     uint      `json:"availabilityId"`
	DoctorID           uint      `json:"doctorId"`
	DayOfWeek          string    `json:"dayOfWeek"`
	AvailableTimeStart string    `json:"availableTimeStart"`
	AvailableTimeEnd   string    `json:"availableTimeEnd"`
	CreatedAt          time.Time `json:"createdAt"`
	CreatedBy          string    `json:"createdBy"`
	ModifiedBy         string    `json:"modifiedBy"`
	ModifiedAt         time.Time `json:"modifiedAt"`
}

func NewDoctorAvailabilityObj() *DoctorAvailabilityObj {
	return &DoctorAvailabilityObj{}
}
