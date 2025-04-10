package requestData

import "time"

type PatientCreationInput struct {
	PatientID      uint      `json:"patientId"`	
	UserID         uint      `json:"userId"`
	FirstName      string    `json:"firstName"` //this is called field tag (the format of object)
	LastName       string    `json:"lastName"`
	DOB            time.Time `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedBy      string    `json:"createdBy"`
	ModifiedBy     string    `json:"modifiedBy"`
	CreatedAt      time.Time `json:"createdAt"`
	ModifiedAt     time.Time `json:"modifiedAt"`
}

func NewPatientCreationInput() *PatientCreationInput {
	return &PatientCreationInput{}
}
