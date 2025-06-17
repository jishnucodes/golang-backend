package requestData

type PrescriptionDosageObj struct {
	PrescriptionDosageID uint   `json:"prescriptionDosageId"`
	PrescriptionDetailID uint   `json:"prescriptionDetailId"`
	Schedule             string `json:"schedule"`
	Dosage               string `json:"dosage"`
	Count                uint    `json:"count"`
	ConsumptionMethod    string `json:"consumptionMethod"`
	Note                 string `json:"note"`
	CreatedAt            string `json:"createdAt"`
	CreatedBy            uint   `json:"createdBy"`
	ModifiedAt           string `json:"modifiedAt"`
	ModifiedBy           uint   `json:"modifiedBy"`
}
