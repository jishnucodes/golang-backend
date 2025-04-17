package builder

import (
	"clinic-management/backend/common"
	"clinic-management/backend/common/requestData"
)

type RoleDTO struct {
	Id         uint   `json:"id"`
	RoleName   string `json:"roleName"`
	Active     uint   `json:"active"`
	CreatedAt  string `json:"createdAt"`
	CreatedBy  string `json:"createdBy"`
	ModifiedAt string `json:"modifiedAt"`
	ModifiedBy string `json:"modifiedBy"`
}

func BuildRoleDTO(role *requestData.RoleObj) *RoleDTO {
	var roleObj RoleDTO

	roleObj.Id = role.Id
	roleObj.RoleName = role.RoleName
	roleObj.Active = role.Active
	roleObj.CreatedAt = role.CreatedAt.Format("2006-01-02 15:04:05")
	roleObj.CreatedBy = role.CreatedBy
	roleObj.ModifiedAt = role.ModifiedAt.Format("2006-01-02 15:04:05")
	roleObj.ModifiedBy = role.ModifiedBy

	return &roleObj
}

func BuildRoleDTOs(rolesData []map[string]interface{}	) []*RoleDTO {
	var roleDTOs []*RoleDTO

	for _, roleMap := range rolesData {
		roleDTO := &RoleDTO{
			Id:         common.ToUint(roleMap["Id"]),
			RoleName:   common.ToString(roleMap["RoleName"]),
			Active:     common.ToUint(roleMap["Active"]),
			CreatedAt:  common.ToString(roleMap["CreatedAt"]),
			CreatedBy:  common.ToString(roleMap["CreatedBy"]),
			ModifiedAt: common.ToString(roleMap["ModifiedAt"]),
			ModifiedBy: common.ToString(roleMap["ModifiedBy"]),
		}
		roleDTOs = append(roleDTOs, roleDTO)
	}
	return roleDTOs
}
