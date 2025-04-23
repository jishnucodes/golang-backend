package requestData

import "time"

type PatientObj struct {
	PatientID      uint      `json:"patientId"`	
	Salutation     string    `json:"salutation"`
	FirstName      string    `json:"firstName"` //this is called field tag (the format of object)
	LastName       string    `json:"lastName"`
	InsuranceID    uint      `json:"insuranceId"`
	BloodGroup     string    `json:"bloodGroup"`
	DOB            time.Time `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedBy      string    `json:"createdBy"`
	ModifiedBy     string    `json:"modifiedBy"`
	CreatedAt      time.Time `json:"createdAt"`
	ModifiedAt     time.Time `json:"modifiedAt"`
}

func NewPatientObj() *PatientObj {
	return &PatientObj{}
}
