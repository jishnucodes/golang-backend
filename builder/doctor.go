package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	// "time"
)

// DoctorDTO represents the structure of the doctor data for API responses
type DoctorDTO struct {
	DoctorID        uint      `json:"doctorId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phoneNumber"`
	MobileNumber    string    `json:"mobileNumber"`
	DepartmentName  string    `json:"departmentName"`
	EmployeeID      uint    `json:"employeeId"`
	DepartmentID    uint    `json:"departmentId"`
	Specialty       string    `json:"specialty"`
	ConsultationFee float64   `json:"consultationFee"`
	CreatedAt       string    `json:"createdAt"`
	CreatedBy       string    `json:"createdBy"`
	ModifiedBy      string    `json:"modifiedBy"`
	ModifiedAt      string    `json:"modifiedAt"`
}

// BuildDoctorDTO constructs and returns a DoctorDTO from doctorData
func BuildDoctorDTO(doctorData *requestData.DoctorObj) *DoctorDTO {
	return &DoctorDTO{
		DoctorID:        doctorData.DoctorID,
		FirstName:       doctorData.FirstName,
		LastName:        doctorData.LastName,
		Email:           doctorData.Email,
		PhoneNumber:     doctorData.PhoneNumber,
		MobileNumber:    doctorData.MobileNumber,
		DepartmentName:  doctorData.DepartmentName,
		EmployeeID:      doctorData.EmployeeID,
		DepartmentID:    doctorData.DepartmentID,
		Specialty:       doctorData.Specialty,
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
			PhoneNumber:     common.ToString(doctorMap["PhoneNumber"]),
			MobileNumber:    common.ToString(doctorMap["MobileNumber"]),
			DepartmentName:  common.ToString(doctorMap["DepartmentName"]),
			EmployeeID:      common.ToUint(doctorMap["EmployeeID"]),
			DepartmentID:    common.ToUint(doctorMap["Department"]),
			Email:           common.ToString(doctorMap["Email"]),
			ConsultationFee: common.ToFloat64(doctorMap["ConsultationFee"]),
			CreatedAt:       common.ToString(doctorMap["CreatedAt"]),
			CreatedBy:       common.ToString(doctorMap["CreatedBy"]),
			ModifiedBy:      common.ToString(doctorMap["ModifiedBy"]),
			ModifiedAt:      common.ToString(doctorMap["ModifiedAt"]),
		}

		doctorDTOs = append(doctorDTOs, doctorDTO)
	}

	return doctorDTOs
}
