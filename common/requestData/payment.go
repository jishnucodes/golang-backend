package requestData

type PaymentObj struct {
	PaymentID           uint    `json:"paymentId"`
	FirstName           string  `json:"firstName"`
	LastName            string  `json:"lastName"`
	TokenNumber         string  `json:"tokenNumber"`
	DoctorID            uint    `json:"doctorId"`
	AppointmentDateTime string  `json:"appointmentDateTime"`
	ConsultationFee     float64 `json:"consultationFee"`
	PaymentDate         string  `json:"paymentDate"`
}

type ConfirmedAppointmentSearchQuery struct {
	AppointmentID *int `form:"appointmentId" json:"appointmentId"`
	ConfirmedBy   *int `form:"confirmedBy" json:"confirmedBy"`
}
