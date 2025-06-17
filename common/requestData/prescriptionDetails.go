package requestData

type PrescriptionDetailsObj struct {
	PrescriptionDetailID uint                    `json:"prescriptionDetailId"`
	ConsultationID       uint                    `json:"consultationId"`
	Medication           string                  `json:"medication"`
	DosageInterval       string                  `json:"dosageInterval"`
	Notes                string                  `json:"notes"`
	Duration             string                  `json:"duration"`
	Frequency            string                  `json:"frequency"`
	Dosage               []PrescriptionDosageObj `json:"dosage"`
	CreatedAt            string                  `json:"createdAt"`
	CreatedBy            uint                    `json:"createdBy"`
	ModifiedAt           string                  `json:"modifiedAt"`
	ModifiedBy           uint                    `json:"modifiedBy"`
}
