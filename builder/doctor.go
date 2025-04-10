package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"time"
)

// DoctorDTO represents the structure of the doctor data for API responses
type DoctorDTO struct {
	DoctorID        uint      `json:"doctorId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Specialty       string    `json:"specialty"`
	ContactNumber   string    `json:"contactNumber"`
	Email           string    `json:"email"`
	ConsultationFee float64   `json:"consultationFee"`
	CreatedAt       time.Time `json:"createdAt"`
	CreatedBy       string    `json:"createdBy"`
	ModifiedBy      string    `json:"modifiedBy"`
	ModifiedAt      time.Time `json:"modifiedAt"`
}

// BuildDoctorDTO constructs and returns a DoctorDTO from doctorData
func BuildDoctorDTO(doctorData *requestData.DoctorObj) *DoctorDTO {
	return &DoctorDTO{
		DoctorID:        doctorData.DoctorID,
		FirstName:       doctorData.FirstName,
		LastName:        doctorData.LastName,
		Specialty:       doctorData.Specialty,
		ContactNumber:   doctorData.ContactNumber,
		Email:           doctorData.Email,
		ConsultationFee: doctorData.ConsultationFee,
		CreatedAt:       doctorData.CreatedAt,
		CreatedBy:       doctorData.CreatedBy,
		ModifiedBy:      doctorData.ModifiedBy,
		ModifiedAt:      doctorData.ModifiedAt,
	}
}

// BuildDoctorDTOs constructs a slice of DoctorDTO from []map[string]interface{}
func BuildDoctorDTOs(doctorsData []map[string]interface{}) []*DoctorDTO {
	var doctorDTOs []*DoctorDTO

	for _, doctorMap := range doctorsData {
		doctorDTO := &DoctorDTO{
			DoctorID:        common.ToUint(doctorMap["DoctorID"]),
			FirstName:       common.ToString(doctorMap["FirstName"]),
			LastName:        common.ToString(doctorMap["LastName"]),
			Specialty:       common.ToString(doctorMap["Specialty"]),
			ContactNumber:   common.ToString(doctorMap["ContactNumber"]),
			Email:           common.ToString(doctorMap["Email"]),
			ConsultationFee: common.ToFloat64(doctorMap["ConsultationFee"]),
			CreatedAt:       common.ParseTime(doctorMap["CreatedAt"]),
			CreatedBy:       common.ToString(doctorMap["CreatedBy"]),
			ModifiedBy:      common.ToString(doctorMap["ModifiedBy"]),
			ModifiedAt:      common.ParseTime(doctorMap["ModifiedAt"]),
		}

		doctorDTOs = append(doctorDTOs, doctorDTO)
	}

	return doctorDTOs
}
