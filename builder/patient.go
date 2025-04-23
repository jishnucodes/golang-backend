package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	// "time"
)

// UserDTO represents the structure of the user data.
type PatientDTO struct {
	PatientID      uint      `json:"patientId"`
	Salutation     string    `json:"salutation"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	InsuranceID    uint      `json:"insuranceId"`
	BloodGroup     string    `json:"bloodGroup"`
	DOB            string `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedAt      string    `json:"createdAt"`
	ModifiedAt     string    `json:"modifiedAt"`
	CreatedBy      string    `json:"createdBy"`
	ModifiedBy     string    `json:"modifiedBy"`
}

// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildPatientDTO(patientData *requestData.PatientObj) *PatientDTO {
	var patientObj PatientDTO

	patientObj.PatientID = patientData.PatientID
	patientObj.Salutation = patientData.Salutation
	patientObj.FirstName = patientData.FirstName
	patientObj.LastName = patientData.LastName
	patientObj.InsuranceID = patientData.InsuranceID
	patientObj.BloodGroup = patientData.BloodGroup
	patientObj.DOB = patientData.DOB.Format("2006-01-02 15:04:05")
	patientObj.Gender = patientData.Gender
	patientObj.MedicalHistory = patientData.MedicalHistory
	patientObj.CreatedAt = patientData.CreatedAt.Format("2006-01-02 15:04:05")
	patientObj.ModifiedAt = patientData.ModifiedAt.Format("2006-01-02 15:04:05")
	patientObj.CreatedBy = patientData.CreatedBy
	patientObj.ModifiedBy = patientData.ModifiedBy

	return &patientObj
}

// BuildPatientDTOs constructs a slice of PatientDTO from []map[string]interface{}
func BuildPatientDTOs(patientsData []map[string]interface{}) []*PatientDTO {
	var patientDTOs []*PatientDTO

	for _, patientMap := range patientsData {
		patientDTO := &PatientDTO{
			PatientID:      common.ToUint(patientMap["PatientID"]),
			Salutation:     common.ToString(patientMap["Salutation"]),
			FirstName:      common.ToString(patientMap["FirstName"]),
			LastName:       common.ToString(patientMap["LastName"]),
			InsuranceID:    common.ToUint(patientMap["InsuranceID"]),
			BloodGroup:     common.ToString(patientMap["BloodGroup"]),
			DOB:            common.ToString(patientMap["DOB"]),
			Gender:         common.ToString(patientMap["Gender"]),
			MedicalHistory: common.ToString(patientMap["MedicalHistory"]),
			CreatedAt:      common.ToString(patientMap["CreatedAt"]),
			ModifiedAt:     common.ToString(patientMap["ModifiedAt"]),
			CreatedBy:      common.ToString(patientMap["CreatedBy"]),
			ModifiedBy:     common.ToString(patientMap["ModifiedBy"]),
		}

		patientDTOs = append(patientDTOs, patientDTO)
	}
	return patientDTOs
}
