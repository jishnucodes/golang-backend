package models

import (
	"time"
)

// CMSDoctorAvailability represents the CMS_DoctorAvailability table
type CMSDoctorAvailability struct {
	AvailabilityID     uint      `gorm:"primaryKey;autoIncrement;column:AvailabilityID"`
	DoctorID           uint      `gorm:"column:DoctorID"`
	DayOfWeek          string    `gorm:"column:DayOfWeek;type:enum('Monday','Tuesday','Wednesday','Thursday','Friday','Saturday','Sunday')"`
	AvailableTimeStart time.Time `gorm:"column:AvailableTimeStart;type:time"`
	AvailableTimeEnd   time.Time `gorm:"column:AvailableTimeEnd;type:time"`
	WeekType           string    `gorm:"column:WeekType"`
	CreatedAt          time.Time `gorm:"column:CreatedAt"`
	CreatedBy          uint    `gorm:"column:CreatedBy"`
	ModifiedBy         uint    `gorm:"column:ModifiedBy"`
	ModifiedAt         time.Time `gorm:"column:ModifiedAt"`

	// Relationship with Doctor
	Doctor CMSDoctor `gorm:"foreignKey:DoctorID;references:DoctorID;constraint:OnDelete:CASCADE"`
}

// TableName overrides the default table name
func (CMSDoctorAvailability) TableName() string {
	return "CMS_DoctorAvailability"
}
