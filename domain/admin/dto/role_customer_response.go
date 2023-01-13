package dto

import (
	"github.com/google/uuid"
)

type RoleAdminResponse struct {
	AdminRolesId uuid.UUID
	AdminRoles   string `json:"admin_roles"`
}
