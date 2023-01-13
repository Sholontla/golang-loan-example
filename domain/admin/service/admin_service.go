package service

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/projects/loans/domain/admin"
	dto "github.com/projects/loans/domain/admin/dto"
	"github.com/projects/loans/middleware"
	"github.com/projects/loans/utils/date_utils"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type AdminRepoService struct {
	repo admin.AdminRepository
}

func (req AdminRepoService) NewAdmin(u admin.Admin) (*dto.AdminResponse, *utilerrors.RestErr) {

	c := admin.Admin{
		AdminId:          uuid.New(),
		AdminFirstName:   u.AdminFirstName,
		AdminLastName:    u.AdminLastName,
		AdminUserName:    u.AdminUserName,
		AdminEmail:       u.AdminEmail,
		AdminPhoneNumber: u.AdminPhoneNumber,
		AdminAccount:     u.AdminAccount,
		AdminPassword:    u.AdminPassword,
		AdminCreatedAt:   date_utils.GetNowString(),
		IsAdmin:          u.IsAdmin,
	}

	newAdmin, err := req.repo.AdminRegistration(c)
	if err != nil {
		return nil, err
	}

	response := newAdmin.ToNewAdminResponseDto()

	return &response, nil
}

func (req AdminRepoService) LoginAdmin(ctx *fiber.Ctx, u admin.AdminLogin) (*dto.AdminResponse, utilerrors.RestErr) {
	c := admin.Admin{
		AdminEmail:    u.AdminEmail,
		AdminPassword: u.AdminPassword,
	}

	newAdmin, err := req.repo.AdminLogin(c)
	if err != nil {
		return nil, err
	}

	isAdmin := strings.Contains(ctx.Path(), "/api/loan/admin")
	var scope string

	if isAdmin {
		scope = "admin"
	} else {
		scope = "customer"
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)

	token, errL := middleware.GenerateJWT(u.AdminEmail, scope)
	if errL != nil {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Invalid Credentials ..."})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expireTime,
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	response := newAdmin.ToNewAdminResponseDto()

	return &response, nil
}

func (req AdminRepoService) AdminLogged(payload string) (*dto.AdminResponse, error) {

	customer, err := req.repo.AdminLoggged(payload)
	if err != nil {
		return nil, err
	}
	response := customer.ToNewAdminResponseDto()
	return &response, nil
}

func (cusHandl AdminRepoService) AdminLogout(c *fiber.Ctx) error {

	nowTime := time.Now()
	expireTime := nowTime.Add(-time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  expireTime,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return nil
}

func (req AdminRepoService) AdminUpdateInfoService(u admin.Admin) (*admin.Admin, error) {

	c := admin.Admin{
		AdminFirstName:   u.AdminFirstName,
		AdminLastName:    u.AdminLastName,
		AdminUserName:    u.AdminUserName,
		AdminEmail:       u.AdminEmail,
		AdminPhoneNumber: u.AdminPhoneNumber,
		AdminUpdatedAt:   date_utils.GetNowString(),
	}

	req.repo.AdminUpdateInfoDao(c)

	return &c, nil
}

func (req AdminRepoService) AdminUpdatePasswordInfoService(u admin.Admin) (*admin.Admin, error) {

	c := admin.Admin{
		AdminPassword: u.AdminPassword,
		AdminEmail:    u.AdminEmail,
	}

	req.repo.AdminUpdatePasswordDao(c)

	return &c, nil
}

func NewAdminService(repo admin.AdminRepository) AdminRepoService {
	return AdminRepoService{repo: repo}
}
