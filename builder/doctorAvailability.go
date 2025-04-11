package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"fmt"
	// "time"
)

// DoctorAvailabilityDTO represents the structure of the doctor availability data for API responses
type DoctorAvailabilityDTO struct {
	AvailabilityID     uint      `json:"availabilityId"`
	DoctorID           uint      `json:"doctorId"`
	DayOfWeek          string    `json:"dayOfWeek"`
	// AvailableTimeStart time.Time `json:"availableTimeStart"`
	// AvailableTimeEnd   time.Time `json:"availableTimeEnd"`
	AvailableTimeStart string `json:"availableTimeStart"`
	AvailableTimeEnd   string `json:"availableTimeEnd"`
	// CreatedAt          time.Time `json:"createdAt"`
	CreatedAt          string `json:"createdAt"`
	CreatedBy          string `json:"createdBy"`
	ModifiedBy         string `json:"modifiedBy"`
	ModifiedAt         string `json:"modifiedAt"`
}

// BuildDoctorAvailabilityDTO constructs and returns a DoctorAvailabilityDTO from availabilityData
func BuildDoctorAvailabilityDTO(availabilityData *requestData.DoctorAvailabilityObj) *DoctorAvailabilityDTO {
	fmt.Println("availabilityData", availabilityData.AvailableTimeStart)
	fmt.Println("availabilityData", availabilityData.AvailableTimeEnd)
	return &DoctorAvailabilityDTO{
		AvailabilityID:     availabilityData.AvailabilityID,
		DoctorID:           availabilityData.DoctorID,
		DayOfWeek:          availabilityData.DayOfWeek,
		AvailableTimeStart: availabilityData.AvailableTimeStart,
		AvailableTimeEnd:   availabilityData.AvailableTimeEnd,
		CreatedAt:          availabilityData.CreatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy:          availabilityData.CreatedBy,
		ModifiedBy:         availabilityData.ModifiedBy,
		ModifiedAt:         availabilityData.ModifiedAt.Format("2006-01-02 15:04:05"),
	}
}

// BuildDoctorAvailabilityDTOs constructs a slice of DoctorAvailabilityDTO from []map[string]interface{}
func BuildDoctorAvailabilityDTOs(availabilitiesData []map[string]interface{}) []*DoctorAvailabilityDTO {
	var availabilityDTOs []*DoctorAvailabilityDTO

	for _, availabilityMap := range availabilitiesData {
		availabilityDTO := &DoctorAvailabilityDTO{
			AvailabilityID:     common.ToUint(availabilityMap["AvailabilityID"]),
			DoctorID:           common.ToUint(availabilityMap["DoctorID"]),
			DayOfWeek:          common.ToString(availabilityMap["DayOfWeek"]),
			AvailableTimeStart: common.ToString(availabilityMap["AvailableTimeStart"]),
			AvailableTimeEnd:   common.ToString(availabilityMap["AvailableTimeEnd"]),
			CreatedAt:          common.ToString(availabilityMap["CreatedAt"]),
			CreatedBy:          common.ToString(availabilityMap["CreatedBy"]),
			ModifiedBy:         common.ToString(availabilityMap["ModifiedBy"]),
			ModifiedAt:         common.ToString(availabilityMap["ModifiedAt"]),
		}

		availabilityDTOs = append(availabilityDTOs, availabilityDTO)
	}

	return availabilityDTOs
}
