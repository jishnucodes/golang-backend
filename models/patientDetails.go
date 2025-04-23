package models

import "time"

// CMSUser represents the CMS_Users table
type CMSPatients struct {
	PatientID        uint      `gorm:"primaryKey;autoIncrement;column:PatientID"`
	Salutation      string    `gorm:"column:Salutation"`
	FirstName     string    `gorm:"column:FirstName"`
	LastName      string    `gorm:"column:LastName"`
	InsuranceID   uint      `gorm:"column:InsuranceID"`
	BloodGroup    string    `gorm:"column:BloodGroup"`
	DOB           time.Time `gorm:"column:DOB"`                     // Corrected: date → time.Time
	Gender        string    `gorm:"column:Gender"`
	MedicalHistory string    `gorm:"column:MedicalHistory"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
	CreatedBy     string    `gorm:"column:CreatedBy"`               // Corrected: varchar → string
	ModifiedAt    time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy    string    `gorm:"column:ModifiedBy"`              // Corrected: varchar → string

	// User           CMSUser      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (CMSPatients) TableName() string {
	return "CMS_Patients"
}
