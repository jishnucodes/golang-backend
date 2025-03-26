package common

import "time"

type PatientCreationInput struct {
	UserID         uint      `json:"userId"`
	FirstName      string    `json:"firstName"` //this is called field tag (the format of object)
	LastName       string    `json:"lastName"`
	DOB            time.Time `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedBy      string    `json:"createdBy"`
	ModifiedBy     string    `json:"modifiedBy"`
}

func NewPatientCreationInput() *PatientCreationInput {
	return &PatientCreationInput{}
}
