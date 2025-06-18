package models

import (
	"time"
)

// MedicineMaster represents a Medicine in the clinic
type CMSMedicineMaster struct {
	MedicineID    uint      `gorm:"column:MedicineID;primaryKey;autoIncrement"`
	MedicineName  string    `gorm:"column:MedicineName;type:varchar(255);not null"`
	GenericName   string    `gorm:"column:GenericName ;type:varchar(100)"`
	BrandName     string    `gorm:"column:BrandName;type:varchar(100)"`
	Composition   string    `gorm:"column:Composition;type:varchar(100)"`
	Form          string    `gorm:"column:Form;type:varchar(100)"`
	Strength      int       `gorm:"column:Strength;type:varchar(50)"`
	Unit          string    `gorm:"column:Unit;type:varchar(50)"`
	Manufacturer  string    `gorm:"column:Manufacturer;type:nvarchar(100)"`
	HsnCode       string    `gorm:"column:HsnCode;type:varchar(20)"`
	GstPercentage float64   `gorm:"column:GstPercentage;type:decimal(5,2)"`
	ReorderLevel  uint      `gorm:"column:ReorderLevel;default:0"`
	Status        uint      `gorm:"column:Status;type:bit"`
	CreatedAt     time.Time `gorm:"column:CreatedAt"`
	CreatedBy     uint      `gorm:"column:CreatedBy;type:int"`
	ModifiedAt    time.Time `gorm:"column:ModifiedAt"`
	ModifiedBy    uint      `gorm:"column:ModifiedBy;type:int"`
}

// TableName specifies the table name for the MedicineMaster model
func (CMSMedicineMaster) TableName() string {
	return "CMS_MedicineMaster"
}
