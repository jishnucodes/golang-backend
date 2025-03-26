package models

import "time"

// CMSUser represents the CMS_Users table
type CMSUser struct {
	UserID        uint      `gorm:"primaryKey;autoIncrement;column:UserID"`
	FirstName     string    `gorm:"column:FirstName"`
	LastName      string    `gorm:"column:LastName"`
	DOB           time.Time `gorm:"column:DOB"`                     // Corrected: date → time.Time
	Gender        string    `gorm:"column:Gender"`
	ContactNumber string    `gorm:"column:ContactNumber"`
	Email         string    `gorm:"column:Email"`
	Address       string    `gorm:"column:Address"`
	Role          string       `gorm:"column:Role"`
	BiometricData []byte    `gorm:"column:BiometricData"`           // Corrected: varbinary → []byte
	PasswordHash  string    `gorm:"column:PasswordHash"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
	CreatedBy     string    `gorm:"column:CreatedBy"`               // Corrected: varchar → string
	ModifiedAt    time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy    string    `gorm:"column:ModifiedBy"`              // Corrected: varchar → string
}

// TableName overrides the default table name
func (CMSUser) TableName() string {
	return "CMS_Users"
}
