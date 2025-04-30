package models

// RoleMaster represents the role master table in the database
type CMSAutoNumber struct {
	Type      uint   `gorm:"column:Type"`
	Prefix    string `gorm:"column:Prefix"`
	Increment uint   `gorm:"column:Increment"`
	Suffix    string `gorm:"column:Suffix"`
}

// TableName specifies the table name for RoleMaster
func (CMSAutoNumber) TableName() string {
	return "CMS_AutoNumber"
}
