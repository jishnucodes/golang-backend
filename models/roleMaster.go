package models

import "time"

// RoleMaster represents the role master table in the database
type CMSRolesMaster struct {
	Id         uint           `gorm:"primaryKey;column:Id"`
	RoleName   string        `gorm:"column:RoleName;type:nvarchar(100);not null"`
	Active     uint          `gorm:"column:Active;type:bit"`
	CreatedAt  time.Time        `gorm:"column:CreatedAt;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreatedBy  string        `gorm:"column:CreatedBy;type:varchar(50)"`
	ModifiedBy string        `gorm:"column:ModifiedBy;type:varchar(50)"`
	ModifiedAt time.Time        `gorm:"column:ModifiedAt;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for RoleMaster
func (CMSRolesMaster) TableName() string {
	return "CMS_Roles_Master"
}
