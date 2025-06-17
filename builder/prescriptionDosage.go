package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type PrescriptionDosageDTO struct {
	PrescriptionDosageID uint   `json:"prescriptionDosageId"`
	PrescriptionDetailID uint   `json:"prescriptionDetailId"`
	Schedule             string `json:"schedule"`
	Dosage               string `json:"dosage"`
	Count                uint   `json:"count"`
	ConsumptionMethod    string `json:"consumptionMethod"`
	Note                 string `json:"note"`
	CreatedAt            string `json:"createdAt"`
	CreatedBy            uint   `json:"createdBy"`
	ModifiedAt           string `json:"modifiedAt"`
	ModifiedBy           uint   `json:"modifiedBy"`
}

func BuildPrescriptionDosageDTO(prescriptionDosageObj *requestData.PrescriptionDosageObj) *PrescriptionDosageDTO {

	var prescriptionDosageDTO PrescriptionDosageDTO

	prescriptionDosageDTO.PrescriptionDosageID = prescriptionDosageObj.PrescriptionDosageID
	prescriptionDosageDTO.PrescriptionDetailID = prescriptionDosageObj.PrescriptionDetailID
	prescriptionDosageDTO.Schedule = prescriptionDosageObj.Schedule
	prescriptionDosageDTO.Dosage = prescriptionDosageObj.Dosage
	prescriptionDosageDTO.Count = prescriptionDosageObj.Count
	prescriptionDosageDTO.ConsumptionMethod = prescriptionDosageObj.ConsumptionMethod
	prescriptionDosageDTO.Note = prescriptionDosageObj.Note
	prescriptionDosageDTO.CreatedAt = prescriptionDosageObj.CreatedAt
	prescriptionDosageDTO.CreatedBy = prescriptionDosageObj.CreatedBy
	prescriptionDosageDTO.ModifiedAt = prescriptionDosageObj.ModifiedAt
	prescriptionDosageDTO.ModifiedBy = prescriptionDosageObj.ModifiedBy

	return &prescriptionDosageDTO

}

func BuildPrescriptionDosageDTOs(prescriptionDosageData []map[string]interface{}) []*PrescriptionDosageDTO {
	var prescriptionDosageDTOs []*PrescriptionDosageDTO

	for _, prescriptionDosageMap := range prescriptionDosageData {
		prescriptionDosageDTO := &PrescriptionDosageDTO{
			PrescriptionDosageID: common.ToUint(prescriptionDosageMap["PrescriptionDosageID"]),
			PrescriptionDetailID: common.ToUint(prescriptionDosageMap["PrescriptionDetailID"]),
			Schedule:             common.ToString(prescriptionDosageMap["Schedule"]),
			Dosage:               common.ToString(prescriptionDosageMap["Dosage"]),
			Count:                common.ToUint(prescriptionDosageMap["Count"]),
			ConsumptionMethod:    common.ToString(prescriptionDosageMap["ConsumptionMethod"]),
			Note:                 common.ToString(prescriptionDosageMap["Note"]),
			CreatedAt:            common.ToString(prescriptionDosageMap["CreatedAt"]),
			ModifiedAt:           common.ToString(prescriptionDosageMap["ModifiedAt"]),
			CreatedBy:            common.ToUint(prescriptionDosageMap["CreatedBy"]),
			ModifiedBy:           common.ToUint(prescriptionDosageMap["ModifiedBy"]),
		}

		prescriptionDosageDTOs = append(prescriptionDosageDTOs, prescriptionDosageDTO)
	}

	return prescriptionDosageDTOs
}
