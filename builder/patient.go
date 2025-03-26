package builder

import (
	"clinic-management/backend/common"
	"time"
)

// UserDTO represents the structure of the user data.
type PatientDTO struct {
	UserID    uint      
	FirstName string    
	LastName  string    
	DOB       time.Time 
	Gender  string   
	MedicalHistory  string 
	CreatedBy  string  
	ModifiedBy  string  
}

// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildPatientDTO(patientData *common.PatientCreationInput) *PatientDTO {
	var patientObj PatientDTO

	patientObj.UserID = patientData.UserID
	patientObj.FirstName = patientData.FirstName
	patientObj.LastName = patientData.LastName
	patientObj.DOB = patientData.DOB
	patientObj.Gender = patientData.Gender
	patientObj.MedicalHistory = patientData.MedicalHistory
	patientObj.CreatedBy = patientData.CreatedBy
	patientObj.ModifiedBy = patientData.ModifiedBy

	return &patientObj;
}
