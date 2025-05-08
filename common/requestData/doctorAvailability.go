package requestData


type DoctorAvailabilityObj struct {
	AvailabilityID     uint      `json:"availabilityId"`
	DoctorID           uint      `json:"doctorId"`
	DayOfWeek          string    `json:"dayOfWeek"`
	AvailableTimeStart string    `json:"availableTimeStart"`
	AvailableTimeEnd   string    `json:"availableTimeEnd"`
	WeekType           string    `json:"weekType"`
	CreatedAt          string    `json:"createdAt"`
	CreatedBy          string    `json:"createdBy"`
	ModifiedBy         string    `json:"modifiedBy"`
	ModifiedAt         string    `json:"modifiedAt"`
}

func NewDoctorAvailabilityObj() *DoctorAvailabilityObj {
	return &DoctorAvailabilityObj{}
}
