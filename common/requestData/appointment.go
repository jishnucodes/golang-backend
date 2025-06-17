package requestData

// import "time"

type AppointmentObj struct {
	AppointmentID       uint   `json:"appointmentId"`
	UserID              uint   `json:"userId"`
	PatientID           uint   `json:"patientId"`
	DoctorID            uint   `json:"doctorId"`
	AppointmentDateTime string `json:"appointmentDateTime"`
	TokenNumber         uint   `json:"tokenNumber"`
	AppointmentType     uint   `json:"appointmentType"`
	Status              uint   `json:"status"`
	CallingCount        uint   `json:"callingCount"`
	PaymentDate         string `json:"paymentDate"`
	ConfirmedBy         uint   `json:"confirmedBy"`
	OrderIndex          uint   `json:"orderIndex"`
	CreatedAt           string `json:"createdAt"`
	CreatedBy           uint   `json:"createdBy"`
	ModifiedAt          string `json:"modifiedAt"`
	ModifiedBy          uint   `json:"modifiedBy"`
}

type AppointmentSearchQuery struct {
	DoctorID  *int    `form:"doctorId" json:"doctorId"`
	InputDate *string `form:"inputDate" json:"inputDate"`
}

type ActiveAppointmentPatientSearchQuery struct {
	Status   *int    `form:"status" json:"status"`
	InputDate *string `form:"inputDate" json:"inputDate"`
}

func NewAppointmentObj() *AppointmentObj {
	return &AppointmentObj{}
}
