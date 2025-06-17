package requestData

type ConsultationMedicationObj struct {
	ConsultationID      uint                     `json:"consultationId"`
	AppointmentID       uint                     `json:"appointmentId"`
	PatientID           uint                   `json:"patientId"`
	DoctorID            uint                   `json:"doctorId"`
	Notes               string                   `json:"notes"`
	NextAppointmentDate string                   `json:"nextAppointmentDate"`
	Prescription        []PrescriptionDetailsObj `json:"prescription"`
	CreatedAt           string                   `json:"createdAt"`
	CreatedBy           uint                     `json:"createdBy"`
	ModifiedAt          string                   `json:"modifiedAt"`
	ModifiedBy          uint                     `json:"modifiedBy"`
}


func NewConsultationMedicationObj() *ConsultationMedicationObj {
	return &ConsultationMedicationObj{}
}
