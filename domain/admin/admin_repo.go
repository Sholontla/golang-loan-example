package admin

import utilerrors "github.com/projects/loans/utils/util_errors"

type AdminRepository interface {
	AdminRegistration(Admin) (*Admin, *utilerrors.RestErr)
	AdminLogin(Admin) (*Admin, utilerrors.RestErr)
	AdminLoggged(payload string) (*Admin, error)
	AdminUpdateInfoDao(Admin) (*Admin, error)
	AdminUpdatePasswordDao(Admin) (*Admin, error)
}
