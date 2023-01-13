package admin

import utilerrors "github.com/projects/loans/utils/util_errors"

type RoleAdminRepository interface {
	RoleAdminRegistration(AdminRoles) (*AdminRoles, *utilerrors.RestErr)
	RoleAdminUpdateInfoDao(AdminRoles) (*AdminRoles, error)
}
