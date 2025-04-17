package requestData

import (
	"time"
)

type RoleObj struct {
	Id        uint      `json:"id"`
	RoleName   string    `json:"roleName"`
	Active     uint      `json:"active"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	ModifiedAt time.Time `json:"modifiedAt"`
	ModifiedBy string    `json:"modifiedBy"`
}

func NewRoleObj() *RoleObj {
	return &RoleObj{}
}
