package models

import "time"

// CMSUser represents the CMS_Users table
type CMSPrescriptionDosage struct {
	PrescriptionDosageID uint      `gorm:"primaryKey;autoIncrement;column:PrescriptionDosageID"`
	PrescriptionDetailID uint      `gorm:"column:PrescriptionDetailID"`
	Schedule             string    `gorm:"column:Schedule"`
	Dosage               string    `gorm:"column:Dosage"`
	Count                int       `gorm:"column:Count"`
	ConsumptionMethod    string    `gorm:"column:ConsumptionMethod"`
	Note                 string    `gorm:"column:Note"`
	CreatedAt            time.Time `gorm:"column:CreatedAt"`
	CreatedBy            uint      `gorm:"column:CreatedBy;type:int"`
	ModifiedAt           time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy           uint      `gorm:"column:ModifiedBy;type:int"`

	PrescriptionDetail CMSPrescriptionDetails `gorm:"foreignKey:PrescriptionDetailID;references:PrescriptionDetailID;constraint:OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (CMSPrescriptionDosage) TableName() string {
	return "CMS_PrescriptionDosage"
}
