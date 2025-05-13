package models

import (
	"time"
)

// UserRole represents the user roles mapping in the database
type CMSUserRoles struct {
	ID         uint           `gorm:"primaryKey;column:Id"`
	RoleID     uint           `gorm:"column:RoleId;not null"`
	UserID     uint           `gorm:"column:UserId;not null"`
	CreatedAt  time.Time      `gorm:"column:CreatedAt;type:timestamp;default:CURRENT_TIMESTAMP"`
	CreatedBy  uint         `gorm:"column:CreatedBy;type:int"`
	ModifiedBy uint         `gorm:"column:ModifiedBy;type:int"`
	ModifiedAt time.Time      `gorm:"column:ModifiedAt;type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	Role       CMSRolesMaster `gorm:"foreignKey:RoleId"`
	User       CMSUser        `gorm:"foreignKey:UserId"`
}

// TableName specifies the table name for UserRole
func (CMSUserRoles) TableName() string {
	return "CMS_UserRoles"
}
