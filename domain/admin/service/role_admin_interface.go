package service

import (
	"github.com/projects/loans/domain/admin"
	dto "github.com/projects/loans/domain/admin/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type IRoleAdminService interface {
	RoleNewAdmin(admin.AdminRoles) (*dto.RoleAdminResponse, *utilerrors.RestErr)
	RoleAdminUpdateInfoService(admin.AdminRoles) (*admin.AdminRoles, error)
}
