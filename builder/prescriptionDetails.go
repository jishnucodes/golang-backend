package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
	"fmt"
)

type PrescriptionDetailsDTO struct {
	PrescriptionDetailID uint                     `json:"prescriptionDetailId"`
	ConsultationID       uint                     `json:"consultationId"`
	Medication           string                   `json:"medication"`
	DosageInterval       string                   `json:"dosageInterval"`
	Notes                string                   `json:"notes"`
	Duration             string                   `json:"duration"`
	Frequency            string                   `json:"frequency"`
	Dosage               []*PrescriptionDosageDTO `json:"dosage"`
	CreatedAt            string                   `json:"createdAt"`
	CreatedBy            uint                     `json:"createdBy"`
	ModifiedAt           string                   `json:"modifiedAt"`
	ModifiedBy           uint                     `json:"modifiedBy"`
}

func BuildPrescriptionDetailsDTO(prescriptionDetailsObj *requestData.PrescriptionDetailsObj) *PrescriptionDetailsDTO {

	var prescriptionDetailsDTO PrescriptionDetailsDTO

	prescriptionDetailsDTO.PrescriptionDetailID = prescriptionDetailsObj.PrescriptionDetailID
	prescriptionDetailsDTO.ConsultationID = prescriptionDetailsObj.ConsultationID
	prescriptionDetailsDTO.Medication = prescriptionDetailsObj.Medication
	prescriptionDetailsDTO.DosageInterval = prescriptionDetailsObj.DosageInterval
	prescriptionDetailsDTO.Notes = prescriptionDetailsObj.Notes
	prescriptionDetailsDTO.Duration = prescriptionDetailsObj.Duration
	prescriptionDetailsDTO.Frequency = prescriptionDetailsObj.Frequency

	var dosageData []map[string]interface{}
	for _, dosage := range prescriptionDetailsObj.Dosage {
		dosageMap := map[string]interface{}{
			"PrescriptionDosageID": dosage.PrescriptionDosageID,
			"PrescriptionDetailID": dosage.PrescriptionDetailID,
			"Schedule":             dosage.Schedule,
			"Dosage":               dosage.Dosage,
			"Count":                dosage.Count,
			"ConsumptionMethod":    dosage.ConsumptionMethod,
			"Note":                 dosage.Note,
			"CreatedAt":            dosage.CreatedAt,
			"CreatedBy":            dosage.CreatedBy,
			"ModifiedBy":           dosage.ModifiedBy,
			"ModifiedAt":           dosage.ModifiedAt,
		}
		dosageData = append(dosageData, dosageMap)
	}
	fmt.Printf("Dosage Data for %v: %+v\n", prescriptionDetailsObj.PrescriptionDetailID, dosageData)

	prescriptionDetailsDTO.Dosage = BuildPrescriptionDosageDTOs(dosageData)
	prescriptionDetailsDTO.CreatedAt = prescriptionDetailsObj.CreatedAt
	prescriptionDetailsDTO.CreatedBy = prescriptionDetailsObj.CreatedBy
	prescriptionDetailsDTO.ModifiedAt = prescriptionDetailsObj.ModifiedAt
	prescriptionDetailsDTO.ModifiedBy = prescriptionDetailsObj.ModifiedBy

	return &prescriptionDetailsDTO

}



func BuildPrescriptionDetailsDTOs(prescriptionDetailsData []map[string]interface{}) []*PrescriptionDetailsDTO {
	fmt.Printf("DEBUG: BuildPrescriptionDetailsDTOs input count: %d\n", len(prescriptionDetailsData))

	var prescriptionDetailsDTOs []*PrescriptionDetailsDTO

	for i, prescriptionDetailsMap := range prescriptionDetailsData {
		fmt.Printf("DEBUG: Processing prescription %d\n", i)

		// Handle dosage conversion
		var dosageDTOs []*PrescriptionDosageDTO
		if rawDosageDetails, ok := prescriptionDetailsMap["Dosage"]; ok && rawDosageDetails != nil {
			fmt.Printf("DEBUG: Raw dosage type: %T\n", rawDosageDetails)

			if dosageData, ok := rawDosageDetails.([]map[string]interface{}); ok {
				fmt.Printf("DEBUG: Found %d dosage items\n", len(dosageData))
				dosageDTOs = BuildPrescriptionDosageDTOs(dosageData)
				fmt.Printf("DEBUG: Created %d dosage DTOs\n", len(dosageDTOs))
			} else if rawDosageSlice, ok := rawDosageDetails.([]interface{}); ok {
				// Case 2: Convert []interface{} to []map[string]interface{}
				fmt.Printf("DEBUG: Found %d dosage items (as []interface{})\n", len(rawDosageSlice))

				var dosageData []map[string]interface{}
				for _, item := range rawDosageSlice {
					if dosageMap, ok := item.(map[string]interface{}); ok {
						dosageData = append(dosageData, dosageMap)
					} else {
						fmt.Printf("DEBUG: Skipping invalid dosage item: %v\n", item)
					}
				}
				dosageDTOs = BuildPrescriptionDosageDTOs(dosageData)

			} else {
				fmt.Printf("DEBUG: Failed to convert dosage data to []map[string]interface{}\n")
			}
		}

		prescriptionDetailsDTO := &PrescriptionDetailsDTO{
			PrescriptionDetailID: common.ToUint(prescriptionDetailsMap["PrescriptionDetailID"]),
			ConsultationID:       common.ToUint(prescriptionDetailsMap["ConsultationID"]),
			Medication:           common.ToString(prescriptionDetailsMap["Medication"]),
			DosageInterval:       common.ToString(prescriptionDetailsMap["DosageInterval"]),
			Notes:                common.ToString(prescriptionDetailsMap["Notes"]),
			Duration:             common.ToString(prescriptionDetailsMap["Duration"]),
			Frequency:            common.ToString(prescriptionDetailsMap["Frequency"]),
			Dosage:               dosageDTOs,
			CreatedAt:            common.ToString(prescriptionDetailsMap["CreatedAt"]),
			ModifiedAt:           common.ToString(prescriptionDetailsMap["ModifiedAt"]),
			CreatedBy:            common.ToUint(prescriptionDetailsMap["CreatedBy"]),
			ModifiedBy:           common.ToUint(prescriptionDetailsMap["ModifiedBy"]),
		}

		fmt.Printf("DEBUG: Final prescription DTO %d: Medication=%s, DosageCount=%d\n",
			i, prescriptionDetailsDTO.Medication, len(prescriptionDetailsDTO.Dosage))

		prescriptionDetailsDTOs = append(prescriptionDetailsDTOs, prescriptionDetailsDTO)
	}

	return prescriptionDetailsDTOs
}
