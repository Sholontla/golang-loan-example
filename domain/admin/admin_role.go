package admin

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/admin/dto"
)

type AdminRoles struct {
	AdminRolesId uuid.UUID
	AdminRoles   string `json:"admin_roles"`
}

func (a AdminRoles) ToNewRoleAdminResponseDto() dto.RoleAdminResponse {
	return dto.RoleAdminResponse{
		AdminRolesId: a.AdminRolesId,
		AdminRoles:   a.AdminRoles,
	}
}
