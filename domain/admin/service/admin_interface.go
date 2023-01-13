package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/admin"
	dto "github.com/projects/loans/domain/admin/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type AdminService interface {
	NewAdmin(admin.Admin) (*dto.AdminResponse, *utilerrors.RestErr)
	LoginAdmin(*fiber.Ctx, admin.AdminLogin) (*dto.AdminResponse, utilerrors.RestErr)
	AdminLogged(payload string) (*dto.AdminResponse, error)
	AdminLogout(*fiber.Ctx) error
	AdminUpdateInfoService(admin.Admin) (*admin.Admin, error)
	AdminUpdatePasswordInfoService(admin.Admin) (*admin.Admin, error)
}
