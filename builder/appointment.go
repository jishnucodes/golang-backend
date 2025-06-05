package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type AppointmentDTO struct {
	AppointmentID       uint   `json:"appointmentId"`
	UserID              uint   `json:"userId"`
	PatientID           uint   `json:"patientId"`
	DoctorID            uint   `json:"doctorId"`
	AppointmentDateTime string `json:"appointmentDateTime"`
	TokenNumber         uint   `json:"tokenNumber"`
	AppointmentType     uint   `json:"appointmentType"`
	Status              uint   `json:"status"`
	CallingCount        uint   `json:"callingCount"`
	OrderIndex          uint   `json:"orderIndex"`
	ConfirmedBy         uint   `json:"confirmedBy"`
	PaymentDate         string `json:"paymentDate"`
	CreatedAt           string `json:"createdAt"`
	CreatedBy           uint   `json:"createdBy"`
	ModifiedAt          string `json:"modifiedAt"`
	ModifiedBy          uint   `json:"modifiedBy"`
}

func BuildAppointmentDTO(appointmentData *requestData.AppointmentObj) *AppointmentDTO {

	var appointmentDTO AppointmentDTO

	appointmentDTO.AppointmentID = appointmentData.AppointmentID
	appointmentDTO.UserID = appointmentData.UserID
	appointmentDTO.PatientID = appointmentData.PatientID
	appointmentDTO.DoctorID = appointmentData.DoctorID
	appointmentDTO.AppointmentDateTime = appointmentData.AppointmentDateTime
	appointmentDTO.TokenNumber = appointmentData.TokenNumber
	appointmentDTO.AppointmentType = appointmentData.AppointmentType
	appointmentDTO.Status = appointmentData.Status
	appointmentDTO.CallingCount = appointmentData.CallingCount
	appointmentDTO.OrderIndex = appointmentData.OrderIndex
	appointmentDTO.ConfirmedBy = appointmentData.ConfirmedBy
	appointmentDTO.PaymentDate = appointmentData.PaymentDate
	appointmentDTO.CreatedAt = appointmentData.CreatedAt
	appointmentDTO.CreatedBy = appointmentData.CreatedBy
	appointmentDTO.ModifiedAt = appointmentData.ModifiedAt
	appointmentDTO.ModifiedBy = appointmentData.ModifiedBy

	return &appointmentDTO

}

func BuildAppointmentDTOs(appointmentData []map[string]interface{}) []*AppointmentDTO {
	var appointmentDTOs []*AppointmentDTO

	for _, appointmentMap := range appointmentData {
		appointmentDTO := &AppointmentDTO{
			AppointmentID:       common.ToUint(appointmentMap["AppointmentID"]),
			UserID:              common.ToUint(appointmentMap["UserID"]),
			PatientID:           common.ToUint(appointmentMap["PatientID"]),
			DoctorID:            common.ToUint(appointmentMap["DoctorID"]),
			AppointmentDateTime: common.ToString(appointmentMap["AppointmentDateTime"]),
			TokenNumber:         common.ToUint(appointmentMap["TokenNumber"]),
			AppointmentType:     common.ToUint(appointmentMap["AppointmentType"]),
			Status:              common.ToUint(appointmentMap["Status"]),
			CallingCount:        common.ToUint(appointmentMap["CallingCount"]),
			OrderIndex:          common.ToUint(appointmentMap["OrderIndex"]),
			PaymentDate:         common.ToString(appointmentMap["PaymentDate"]),
			CreatedAt:           common.ToString(appointmentMap["CreatedAt"]),
			ModifiedAt:          common.ToString(appointmentMap["ModifiedAt"]),
			CreatedBy:           common.ToUint(appointmentMap["CreatedBy"]),
			ModifiedBy:          common.ToUint(appointmentMap["ModifiedBy"]),
		}

		appointmentDTOs = append(appointmentDTOs, appointmentDTO)
	}

	return appointmentDTOs
}
