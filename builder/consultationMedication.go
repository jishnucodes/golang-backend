package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"fmt"
)

type ConsultationMedicationDTO struct {
	ConsultationID      uint                      `json:"consultationId"`
	AppointmentID       uint                      `json:"appointmentId"`
	PatientID           uint                      `json:"patientId"`
	DoctorID            uint                      `json:"doctorId"`
	Notes               string                    `json:"notes"`
	NextAppointmentDate string                    `json:"nextAppointmentDate"`
	Prescription        []*PrescriptionDetailsDTO `json:"prescription"`
	CreatedAt           string                    `json:"createdAt"`
	CreatedBy           uint                      `json:"createdBy"`
	ModifiedAt          string                    `json:"modifiedAt"`
	ModifiedBy          uint                      `json:"modifiedBy"`
}

// func BuildConsultationMedicationDTO(consultationDetailsObj *requestData.ConsultationMedicationObj) *ConsultationMedicationDTO {

// 	var consultationMedicationDTO ConsultationMedicationDTO

// 	consultationMedicationDTO.ConsultationID = consultationDetailsObj.ConsultationID
// 	consultationMedicationDTO.AppointmentID = consultationDetailsObj.AppointmentID
// 	consultationMedicationDTO.PatientID = consultationDetailsObj.PatientID
// 	consultationMedicationDTO.DoctorID = consultationDetailsObj.DoctorID
// 	consultationMedicationDTO.Notes = consultationDetailsObj.Notes
// 	consultationMedicationDTO.NextAppointmentDate = consultationDetailsObj.NextAppointmentDate

// 	var prescriptionDetailsData []map[string]interface{}
// 	for _, prescriptionDetails := range consultationDetailsObj.Prescription {

// 		// // Step 1: Convert dosage struct to []map[string]interface{}
// 		var dosageData []map[string]interface{}
// 		for _, d := range prescriptionDetails.Dosage {
// 			dosageMap := map[string]interface{}{
// 				"PrescriptionDosageID": d.PrescriptionDosageID,
// 				"PrescriptionDetailID": d.PrescriptionDetailID,
// 				"Schedule":             d.Schedule,
// 				"Dosage":               d.Dosage,
// 				"Count":                d.Count,
// 				"ConsumptionMethod":    d.ConsumptionMethod,
// 				"Note":                 d.Note,
// 				"CreatedAt":            d.CreatedAt,
// 				"CreatedBy":            d.CreatedBy,
// 				"ModifiedBy":           d.ModifiedBy,
// 				"ModifiedAt":           d.ModifiedAt,
// 			}

// 			dosageData = append(dosageData, dosageMap)
// 		}

// 		// Step 2: Include dosageData inside prescriptionDetailsMap
// 		prescriptionDetailsMap := map[string]interface{}{
// 			"PrescriptionDetailID": prescriptionDetails.PrescriptionDetailID,
// 			"ConsultationID":       prescriptionDetails.ConsultationID,
// 			"Medication":           prescriptionDetails.Medication,
// 			"DosageInterval":       prescriptionDetails.DosageInterval,
// 			"Notes":                prescriptionDetails.Notes,
// 			"Duration":             prescriptionDetails.Duration,
// 			"Frequency":            prescriptionDetails.Frequency,
// 			"Dosage":               dosageData, 
// 			"CreatedAt":            prescriptionDetails.CreatedAt,
// 			"CreatedBy":            prescriptionDetails.CreatedBy,
// 			"ModifiedBy":           prescriptionDetails.ModifiedBy,
// 			"ModifiedAt":           prescriptionDetails.ModifiedAt,
// 		}

// 		prescriptionDetailsData = append(prescriptionDetailsData, prescriptionDetailsMap)
// 	}

// 	consultationMedicationDTO.Prescription = BuildPrescriptionDetailsDTOs(prescriptionDetailsData)
// 	consultationMedicationDTO.CreatedAt = consultationDetailsObj.CreatedAt
// 	consultationMedicationDTO.CreatedBy = consultationDetailsObj.CreatedBy
// 	consultationMedicationDTO.ModifiedAt = consultationDetailsObj.ModifiedAt
// 	consultationMedicationDTO.ModifiedBy = consultationDetailsObj.ModifiedBy

// 	return &consultationMedicationDTO

// }


func BuildConsultationMedicationDTO(consultationDetailsObj *requestData.ConsultationMedicationObj) *ConsultationMedicationDTO {
	// Debug: Print input object
	fmt.Printf("DEBUG: Input consultationDetailsObj: %+v\n", consultationDetailsObj)
	
	var consultationMedicationDTO ConsultationMedicationDTO

	consultationMedicationDTO.ConsultationID = consultationDetailsObj.ConsultationID
	consultationMedicationDTO.AppointmentID = consultationDetailsObj.AppointmentID
	consultationMedicationDTO.PatientID = consultationDetailsObj.PatientID
	consultationMedicationDTO.DoctorID = consultationDetailsObj.DoctorID
	consultationMedicationDTO.Notes = consultationDetailsObj.Notes
	consultationMedicationDTO.NextAppointmentDate = consultationDetailsObj.NextAppointmentDate

	// Debug: Print basic fields copied
	fmt.Printf("DEBUG: Basic fields - ConsultationID: %v, PatientID: %v, DoctorID: %v\n", 
		consultationMedicationDTO.ConsultationID, consultationMedicationDTO.PatientID, consultationMedicationDTO.DoctorID)

	var prescriptionDetailsData []map[string]interface{}
	
	// Debug: Print prescription count
	fmt.Printf("DEBUG: Processing %d prescriptions\n", len(consultationDetailsObj.Prescription))
	
	for i, prescriptionDetails := range consultationDetailsObj.Prescription {
		// Debug: Print current prescription being processed
		fmt.Printf("DEBUG: Processing prescription %d: %+v\n", i, prescriptionDetails)

		// Step 1: Convert dosage struct to []map[string]interface{}
		var dosageData []map[string]interface{}
		
		// Debug: Print dosage count for current prescription
		fmt.Printf("DEBUG: Prescription %d has %d dosages\n", i, len(prescriptionDetails.Dosage))
		
		for j, d := range prescriptionDetails.Dosage {
			// Debug: Print current dosage being processed
			fmt.Printf("DEBUG: Processing dosage %d for prescription %d: %+v\n", j, i, d)
			
			dosageMap := map[string]interface{}{
				"PrescriptionDosageID": d.PrescriptionDosageID,
				"PrescriptionDetailID": d.PrescriptionDetailID,
				"Schedule":             d.Schedule,
				"Dosage":               d.Dosage,
				"Count":                d.Count,
				"ConsumptionMethod":    d.ConsumptionMethod,
				"Note":                 d.Note,
				"CreatedAt":            d.CreatedAt,
				"CreatedBy":            d.CreatedBy,
				"ModifiedBy":           d.ModifiedBy,
				"ModifiedAt":           d.ModifiedAt,
			}

			dosageData = append(dosageData, dosageMap)
		}

		// Debug: Print processed dosage data
		fmt.Printf("DEBUG: Processed dosageData for prescription %d: %+v\n", i, dosageData)

		// Step 2: Include dosageData inside prescriptionDetailsMap
		prescriptionDetailsMap := map[string]interface{}{
			"PrescriptionDetailID": prescriptionDetails.PrescriptionDetailID,
			"ConsultationID":       prescriptionDetails.ConsultationID,
			"Medication":           prescriptionDetails.Medication,
			"DosageInterval":       prescriptionDetails.DosageInterval,
			"Notes":                prescriptionDetails.Notes,
			"Duration":             prescriptionDetails.Duration,
			"Frequency":            prescriptionDetails.Frequency,
			"Dosage":               dosageData, 
			"CreatedAt":            prescriptionDetails.CreatedAt,
			"CreatedBy":            prescriptionDetails.CreatedBy,
			"ModifiedBy":           prescriptionDetails.ModifiedBy,
			"ModifiedAt":           prescriptionDetails.ModifiedAt,
		}

		// Debug: Print prescription details map
		fmt.Printf("DEBUG: Created prescriptionDetailsMap for prescription %d: %+v\n", i, prescriptionDetailsMap)

		prescriptionDetailsData = append(prescriptionDetailsData, prescriptionDetailsMap)
	}

	// Debug: Print final prescription details data before building DTOs
	fmt.Printf("DEBUG: Final prescriptionDetailsData before BuildPrescriptionDetailsDTOs: %+v\n", prescriptionDetailsData)

	consultationMedicationDTO.Prescription = BuildPrescriptionDetailsDTOs(prescriptionDetailsData)
	
	// Debug: Print result of BuildPrescriptionDetailsDTOs
	fmt.Printf("DEBUG: Result from BuildPrescriptionDetailsDTOs: %+v\n", consultationMedicationDTO.Prescription)
	
	consultationMedicationDTO.CreatedAt = consultationDetailsObj.CreatedAt
	consultationMedicationDTO.CreatedBy = consultationDetailsObj.CreatedBy
	consultationMedicationDTO.ModifiedAt = consultationDetailsObj.ModifiedAt
	consultationMedicationDTO.ModifiedBy = consultationDetailsObj.ModifiedBy

	// Debug: Print final DTO before return
	fmt.Printf("DEBUG: Final consultationMedicationDTO: %+v\n", consultationMedicationDTO)

	return &consultationMedicationDTO
}
func BuildConsultationMedicationDTOs(consultationMedicationData []map[string]interface{}) []*ConsultationMedicationDTO {
	var consultationMedicationDTOs []*ConsultationMedicationDTO

	for _, consultationMedicationMap := range consultationMedicationData {

		consultationMedicationDTO := &ConsultationMedicationDTO{
			ConsultationID:      common.ToUint(consultationMedicationMap["ConsultationID"]),
			AppointmentID:       common.ToUint(consultationMedicationMap["AppointmentID"]),
			PatientID:           common.ToUint(consultationMedicationMap["PatientID"]),
			DoctorID:            common.ToUint(consultationMedicationMap["DoctorID"]),
			Notes:               common.ToString(consultationMedicationMap["Notes"]),
			NextAppointmentDate: common.ToString(consultationMedicationMap["NextAppointmentDate"]),

			Prescription: func() []*PrescriptionDetailsDTO {
				if rawPrescriptionDetails, ok := consultationMedicationMap["Prescription"]; ok && rawPrescriptionDetails != nil {
					if prescriptionDetailsSlice, ok := rawPrescriptionDetails.([]interface{}); ok {
						var prescriptionDetailsData []map[string]interface{}
						for _, a := range prescriptionDetailsSlice {
							if prescriptionDetailsMap, ok := a.(map[string]interface{}); ok {
								prescriptionDetailsData = append(prescriptionDetailsData, prescriptionDetailsMap)
							}
						}
						return BuildPrescriptionDetailsDTOs(prescriptionDetailsData)
					}
				}
				return nil
			}(),
			CreatedAt:  common.ToString(consultationMedicationMap["CreatedAt"]),
			ModifiedAt: common.ToString(consultationMedicationMap["ModifiedAt"]),
			CreatedBy:  common.ToUint(consultationMedicationMap["CreatedBy"]),
			ModifiedBy: common.ToUint(consultationMedicationMap["ModifiedBy"]),
		}

		consultationMedicationDTOs = append(consultationMedicationDTOs, consultationMedicationDTO)
	}

	return consultationMedicationDTOs
}
