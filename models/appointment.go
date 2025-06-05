package models

import "time"

// CMSUser represents the CMS_Users table
type CMSAppointments struct {
	AppointmentID       uint      `gorm:"primaryKey;autoIncrement;column:AppointmentID"`
	UserID              uint      `gorm:"column:UserID"`
	PatientID           uint      `gorm:"column:PatientID"`
	DoctorID            uint      `gorm:"column:DoctorID"`
	AppointmentDateTime time.Time `gorm:"column:AppointmentDateTime"`
	TokenNumber         uint      `gorm:"column:TokenNumber"`
	AppointmentType     uint      `gorm:"column:AppointmentType"`
	Status              uint      `gorm:"column:Status"`
	CallingCount        uint      `gorm:"column:CallingCount"`
	OrderIndex          uint      `gorm:"column:OrderIndex"`
	CreatedAt           time.Time `gorm:"column:CreatedAt"`
	CreatedBy           uint      `gorm:"column:CreatedBy;type:int"`
	ModifiedAt          time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy          uint      `gorm:"column:ModifiedBy;type:int"`

	User           CMSUser      `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Patient        CMSPatients  `gorm:"foreignKey:PatientID;references:PatientID;constraint:OnDelete:CASCADE"`
	Doctor 		   CMSDoctor	`gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (CMSAppointments) TableName() string {
	return "CMS_Appointments"
}
