package models

import "time"

// CMSUser represents the CMS_Users table
type CMSConsultationHeader struct {
	ConsultationID      uint      `gorm:"primaryKey;autoIncrement;column:ConsultationID"`
	AppointmentID       uint      `gorm:"column:AppointmentID"`
	PatientID           uint      `gorm:"column:PatientID"`
	DoctorID            uint      `gorm:"column:DoctorID"`
	Notes               string    `gorm:"column:notes"`
	NextAppointmentDate time.Time `gorm:"column:NextAppointmentDate"`
	CreatedAt           time.Time `gorm:"column:CreatedAt"`
	CreatedBy           uint      `gorm:"column:CreatedBy;type:int"`
	ModifiedAt          time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy          uint      `gorm:"column:ModifiedBy;type:int"`

	Patient CMSPatients `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE"`
	Doctor  CMSDoctor   `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (CMSConsultationHeader) TableName() string {
	return "CMS_ConsultationHeader"
}
