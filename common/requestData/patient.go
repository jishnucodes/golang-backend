package requestData

// import "time"

type PatientObj struct {
	PatientID      uint      `json:"patientId"`	
	PatientCode    string    `json:"patientCode"`
	Salutation     string    `json:"salutation"`
	FirstName      string    `json:"firstName"` //this is called field tag (the format of object)
	LastName       string    `json:"lastName"`
	InsuranceID    uint      `json:"insuranceId"`
	BloodGroup     string    `json:"bloodGroup"`
	DOB            string    `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	MobileNumber   string    `json:"mobileNumber"`
	ContactName    string    `json:"contactName"`
	Relation       string    `json:"relation"`
	Type           uint      `json:"type"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedBy      uint    `json:"createdBy"`
	ModifiedBy     uint    `json:"modifiedBy"`
	CreatedAt      string    `json:"createdAt"`
	ModifiedAt     string    `json:"modifiedAt"`
}

func NewPatientObj() *PatientObj {
	return &PatientObj{}
}
