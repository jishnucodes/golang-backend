package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"encoding/json"
	"fmt"
	"log"
	// "time"
)

// DoctorDTO represents the structure of the doctor data for API responses
type DoctorDTO struct {
	DoctorID        uint                     `json:"doctorId"`
	EmployeeID      uint                     `json:"employeeId"`
	Specialty       string                   `json:"specialty"`
	ConsultationFee float64                  `json:"consultationFee"`
	Employee        EmployeeDTO              `json:"employee"`
	Availabilities  []*DoctorAvailabilityDTO `json:"availabilities"`
	CreatedAt       string                   `json:"createdAt"`
	CreatedBy       string                   `json:"createdBy"`
	ModifiedBy      string                   `json:"modifiedBy"`
	ModifiedAt      string                   `json:"modifiedAt"`
}

// BuildDoctorDTO constructs and returns a DoctorDTO from doctorData
func BuildDoctorDTO(doctorData *requestData.DoctorObj) *DoctorDTO {
	var doctorDTO DoctorDTO

	doctorDTO.DoctorID = doctorData.DoctorID
	doctorDTO.EmployeeID = doctorData.EmployeeID
	doctorDTO.Specialty = doctorData.Specialty
	doctorDTO.ConsultationFee = doctorData.ConsultationFee
	doctorDTO.Employee = *BuildEmployeeDTO(&doctorData.Employee)

	var availabilitiesData []map[string]interface{}
	for _, availability := range doctorData.Availabilities {
		availabilityMap := map[string]interface{}{
			"AvailabilityID":     availability.AvailabilityID,
			"DoctorID":           availability.DoctorID,
			"DayOfWeek":          availability.DayOfWeek,
			"AvailableTimeStart": availability.AvailableTimeStart,
			"AvailableTimeEnd":   availability.AvailableTimeEnd,
			"WeekType":           availability.WeekType,
			"CreatedAt":          availability.CreatedAt,
			"CreatedBy":          availability.CreatedBy,
			"ModifiedBy":         availability.ModifiedBy,
			"ModifiedAt":         availability.ModifiedAt,
		}
		availabilitiesData = append(availabilitiesData, availabilityMap)
	}
	doctorDTO.Availabilities = BuildDoctorAvailabilityDTOs(availabilitiesData)
	doctorDTO.CreatedAt = doctorData.CreatedAt
	doctorDTO.CreatedBy = doctorData.CreatedBy
	doctorDTO.ModifiedBy = doctorData.ModifiedBy
	doctorDTO.ModifiedAt = doctorData.ModifiedAt

	return &doctorDTO
}

// BuildDoctorDTOs constructs a slice of DoctorDTO from []map[string]interface{}
func BuildDoctorDTOs(doctorsData []map[string]interface{}) []*DoctorDTO {
	var doctorDTOs []*DoctorDTO

	for _, doctorMap := range doctorsData {
		// Step 1: Unmarshal Employee JSON string
		var employeeObj requestData.EmployeeObj
		if empStr, ok := doctorMap["Employee"].(string); ok {
			err := json.Unmarshal([]byte(empStr), &employeeObj)
			if err != nil {
				log.Printf("Failed to unmarshal Employee JSON: %v", err)
				continue
			}
		} else {
			log.Printf("Employee field not a valid string")
			continue
		}

		fmt.Println("employeeObj", employeeObj)



		// Step 2: Construct DTO after processing fields
		doctorDTO := &DoctorDTO{
			DoctorID:        common.ToUint(doctorMap["DoctorID"]),
			EmployeeID:      common.ToUint(doctorMap["EmployeeID"]),
			Specialty:       common.ToString(doctorMap["Specialty"]),
			ConsultationFee: common.ToFloat64(doctorMap["ConsultationFee"]),
			Employee:        *BuildEmployeeDTO(&employeeObj),
			// Convert []interface{} to []map[string]interface{} for Availabilities
			Availabilities: func() []*DoctorAvailabilityDTO {
				if rawAvailabilities, ok := doctorMap["Availabilities"]; ok && rawAvailabilities != nil {
					if avlSlice, ok := rawAvailabilities.([]interface{}); ok {
						var avlData []map[string]interface{}
						for _, a := range avlSlice {
							if avlMap, ok := a.(map[string]interface{}); ok {
								avlData = append(avlData, avlMap)
							}
						}
						return BuildDoctorAvailabilityDTOs(avlData)
					}
				}
				return nil
			}(),
			CreatedAt:       common.ToString(doctorMap["CreatedAt"]),
			CreatedBy:       common.ToString(doctorMap["CreatedBy"]),
			ModifiedBy:      common.ToString(doctorMap["ModifiedBy"]),
			ModifiedAt:      common.ToString(doctorMap["ModifiedAt"]),
		}

		doctorDTOs = append(doctorDTOs, doctorDTO)
	}

	return doctorDTOs
}

