package models

import "time"

// CMSUser represents the CMS_Users table
type CMSPrescriptionDetails struct {
	PrescriptionDetailID uint      `gorm:"primaryKey;autoIncrement;column:PrescriptionDetailID"`
	ConsultationID       uint      `gorm:"column:ConsultationID"`
	Medication           string    `gorm:"column:Medication"`
	DosageInterval       string    `gorm:"column:DosageInterval"`
	Notes                string    `gorm:"column:notes"`
	Duration             string    `gorm:"column:Duration"`
	Frequency            string    `gorm:"column:Duration"`
	CreatedAt            time.Time `gorm:"column:CreatedAt"`
	CreatedBy            uint      `gorm:"column:CreatedBy;type:int"`
	ModifiedAt           time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy           uint      `gorm:"column:ModifiedBy;type:int"`

	Consultation CMSConsultationHeader `gorm:"foreignKey:ConsultationID;references:ConsultationID;constraint:OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (CMSPrescriptionDetails) TableName() string {
	return "CMS_PrescriptionDetails"
}
