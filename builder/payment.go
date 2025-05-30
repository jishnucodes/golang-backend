package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type PaymentDTO struct {
	PaymentID           uint    `json:"paymentId"`
	FirstName           string  `json:"firstName"`
	LastName            string  `json:"lastName"`
	TokenNumber         string    `json:"tokenNumber"`
	DoctorID            uint    `json:"doctorId"`
	AppointmentDateTime string  `json:"appointmentDateTime"`
	ConsultationFee     float64 `json:"consultationFee"`
	PaymentDate         string  `json:"paymentDate"`
}

func BuildPaymentDTO(paymentData *requestData.PaymentObj) *PaymentDTO {
	var paymentObj PaymentDTO	

	paymentObj.PaymentID = paymentData.PaymentID
	paymentObj.FirstName = paymentData.FirstName
	paymentObj.LastName = paymentData.LastName
	paymentObj.TokenNumber = paymentData.TokenNumber
	paymentObj.DoctorID = paymentData.DoctorID
	paymentObj.AppointmentDateTime = paymentData.AppointmentDateTime
	paymentObj.ConsultationFee = paymentData.ConsultationFee
	paymentObj.PaymentDate = paymentData.PaymentDate
	return &paymentObj
}



func BuildPaymentDTOs(paymentsData []map[string]interface{}) []*PaymentDTO {
	var paymentDTOs []*PaymentDTO

	for _, paymentMap := range paymentsData {
		paymentDTO := &PaymentDTO{
			PaymentID:           common.ToUint(paymentMap["PaymentID"]),
			FirstName:           common.ToString(paymentMap["FirstName"]),
			LastName:			common.ToString(paymentMap["LastName"]),
			TokenNumber:         common.ToString(paymentMap["TokenNumber"]),
			DoctorID:            common.ToUint(paymentMap["DoctorID"]),
			AppointmentDateTime: common.ToString(paymentMap["AppointmentDateTime"]),
			ConsultationFee:     common.ToFloat64(paymentMap["ConsultationFee"]),
			PaymentDate:         common.ToString(paymentMap["PaymentDate"]),
		}

		paymentDTOs = append(paymentDTOs, paymentDTO)
	}
	return paymentDTOs

}