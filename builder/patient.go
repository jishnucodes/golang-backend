package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"time"
)

// UserDTO represents the structure of the user data.
type PatientDTO struct {
	PatientID      uint      `json:"patientId"`
	UserID         uint      `json:"userId"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	DOB            time.Time `json:"dateOfBirth"`
	Gender         string    `json:"gender"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedAt      time.Time `json:"createdAt"`
	ModifiedAt     time.Time `json:"modifiedAt"`
	CreatedBy      string    `json:"createdBy"`
	ModifiedBy     string    `json:"modifiedBy"`
}

// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildPatientDTO(patientData *requestData.PatientCreationInput) *PatientDTO {
	var patientObj PatientDTO

	patientObj.PatientID = patientData.PatientID
	patientObj.UserID = patientData.UserID
	patientObj.FirstName = patientData.FirstName
	patientObj.LastName = patientData.LastName
	patientObj.DOB = patientData.DOB
	patientObj.Gender = patientData.Gender
	patientObj.MedicalHistory = patientData.MedicalHistory
	patientObj.CreatedAt = patientData.CreatedAt
	patientObj.ModifiedAt = patientData.ModifiedAt
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
			UserID:         common.ToUint(patientMap["UserID"]),
			FirstName:      common.ToString(patientMap["FirstName"]),
			LastName:       common.ToString(patientMap["LastName"]),
			DOB:            common.ParseTime(patientMap["DOB"]),
			Gender:         common.ToString(patientMap["Gender"]),
			MedicalHistory: common.ToString(patientMap["MedicalHistory"]),
			CreatedAt:      common.ParseTime(patientMap["CreatedAt"]),
			ModifiedAt:     common.ParseTime(patientMap["ModifiedAt"]),
			CreatedBy:      common.ToString(patientMap["CreatedBy"]),
			ModifiedBy:     common.ToString(patientMap["ModifiedBy"]),
		}

		patientDTOs = append(patientDTOs, patientDTO)
	}
	return patientDTOs
}
