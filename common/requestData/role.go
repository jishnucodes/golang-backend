package requestData

import (
	"time"
)

type RoleObj struct {
	Id        uint      `json:"id"`
	RoleName   string    `json:"roleName"`
	Active     uint      `json:"active"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  uint    `json:"createdBy"`
	ModifiedAt time.Time `json:"modifiedAt"`
	ModifiedBy uint    `json:"modifiedBy"`
}

func NewRoleObj() *RoleObj {
	return &RoleObj{}
}
