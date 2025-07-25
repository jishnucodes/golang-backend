package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	// "time"
)

// UserDTO represents the structure of the user data.
type PatientDTO struct {
	PatientID      uint      `json:"patientId"`
	PatientCode    string    `json:"patientCode"`
	Salutation     string    `json:"salutation"`
	FirstName      string    `json:"firstName"`
	LastName       string    `json:"lastName"`
	InsuranceID    uint      `json:"insuranceId"`
	BloodGroup     string    `json:"bloodGroup"`
	DOB            string    `json:"dateOfBirth"`
	MobileNumber   string    `json:"mobileNumber"`
	ContactName    string    `json:"contactName"`
	Relation       string    `json:"relation"`
	Type           uint      `json:"type"`
	Gender         string    `json:"gender"`
	MedicalHistory string    `json:"medicalHistory"`
	CreatedAt      string    `json:"createdAt"`
	ModifiedAt     string    `json:"modifiedAt"`
	CreatedBy      uint    `json:"createdBy"`
	ModifiedBy     uint    `json:"modifiedBy"`
}

// BuildUserDTO constructs and returns a UserDTO from userData.
func BuildPatientDTO(patientData *requestData.PatientObj) *PatientDTO {
	var patientObj PatientDTO

	patientObj.PatientID = patientData.PatientID
	patientObj.PatientCode = patientData.PatientCode
	patientObj.Salutation = patientData.Salutation
	patientObj.FirstName = patientData.FirstName
	patientObj.LastName = patientData.LastName
	patientObj.InsuranceID = patientData.InsuranceID
	patientObj.BloodGroup = patientData.BloodGroup
	patientObj.DOB = patientData.DOB
	patientObj.MobileNumber = patientData.MobileNumber
	patientObj.ContactName = patientData.ContactName
	patientObj.Relation = patientData.Relation
	patientObj.Type = patientData.Type
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
			PatientCode:    common.ToString(patientMap["PatientCode"]),
			Salutation:     common.ToString(patientMap["Salutation"]),
			FirstName:      common.ToString(patientMap["FirstName"]),
			LastName:       common.ToString(patientMap["LastName"]),
			InsuranceID:    common.ToUint(patientMap["InsuranceID"]),
			BloodGroup:     common.ToString(patientMap["BloodGroup"]),
			DOB:            common.ToString(patientMap["DOB"]),
			MobileNumber:   common.ToString(patientMap["MobileNumber"]),
			ContactName:    common.ToString(patientMap["ContactName"]),
			Relation:       common.ToString(patientMap["Relation"]),
			Type:           common.ToUint(patientMap["Type"]),
			Gender:         common.ToString(patientMap["Gender"]),
			MedicalHistory: common.ToString(patientMap["MedicalHistory"]),
			CreatedAt:      common.ToString(patientMap["CreatedAt"]),
			ModifiedAt:     common.ToString(patientMap["ModifiedAt"]),
			CreatedBy:      common.ToUint(patientMap["CreatedBy"]),
			ModifiedBy:     common.ToUint(patientMap["ModifiedBy"]),
		}

		patientDTOs = append(patientDTOs, patientDTO)
	}
	return patientDTOs
}
