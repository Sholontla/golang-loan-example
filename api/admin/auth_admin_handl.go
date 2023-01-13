package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/admin"
	"github.com/projects/loans/domain/admin/service"
	"github.com/projects/loans/middleware"
	utilerrors "github.com/projects/loans/utils/util_errors"
	"github.com/valyala/fasthttp"
)

type AdminHandler struct {
	service service.AdminService
}

func (cusHandl AdminHandler) AdiminRegistrationHandler(ctx *fiber.Ctx) error {

	var request map[string]string

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	if request["admin_password"] != request["admin_password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": "password not match ...."})
	}

	cu := admin.Admin{
		AdminFirstName:   request["admin_first_name"],
		AdminLastName:    request["admin_last_name"],
		AdminUserName:    request["admin_user_name"],
		AdminEmail:       request["admin_email"],
		AdminPhoneNumber: request["admin_phone_number"],
		AdminAccount:     request["admin_account"],
		AdminPassword:    request["admin_password"],
		IsAdmin:          strings.Contains(ctx.Path(), "/api/loan/admin"),
	}

	service, err := cusHandl.service.NewAdmin(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl AdminHandler) LoginAdmin(ctx *fiber.Ctx) error {
	var cu admin.AdminLogin
	if err := ctx.BodyParser(&cu); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	service, err := cusHandl.service.LoginAdmin(ctx, cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}

	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl AdminHandler) AdminLoggedHandler(ctx *fiber.Ctx) error {

	payLoad, err := middleware.GetUserLogin(ctx)

	if err != nil {
		return err
	}

	service, err := cusHandl.service.AdminLogged(payLoad)
	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(fasthttp.StatusOK).JSON(&fiber.Map{
		"message": service})

}
func (cusHandl AdminHandler) AdminLogoutHandler(c *fiber.Ctx) error {

	cusHandl.service.AdminLogout(c)
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success Logout"})

}

func (cusHandl AdminHandler) AdminUpdateInfoHandler(ctx *fiber.Ctx) error {
	var request admin.AdminRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}

	cu := admin.Admin{
		AdminFirstName:   request.AdminFirstName,
		AdminLastName:    request.AdminLastName,
		AdminUserName:    request.AdminUserName,
		AdminEmail:       payLoad,
		AdminPhoneNumber: request.AdminPhoneNumber,
	}

	service, err := cusHandl.service.AdminUpdateInfoService(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})

}

func (cusHandl AdminHandler) AdminUpdatePasswordInfoHandler(ctx *fiber.Ctx) error {
	var request map[string]string

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if request["admin_password"] != request["admin_password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": "password not match ...."})
	}
	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}
	cu := admin.Admin{
		AdminPassword: request["admin_password"],
		AdminEmail:    payLoad,
	}

	service, err := cusHandl.service.AdminUpdatePasswordInfoService(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})
}
