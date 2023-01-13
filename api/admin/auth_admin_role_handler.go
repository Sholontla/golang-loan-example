package admin

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/admin"
	"github.com/projects/loans/domain/admin/service"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type RoleAdminHandler struct {
	service service.IRoleAdminService
}

func (cusHandl RoleAdminHandler) RoleAdiminRegistrationHandler(ctx *fiber.Ctx) error {

	var request admin.AdminRoles

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	cu := admin.AdminRoles{
		AdminRoles: request.AdminRoles,
	}
	service, err := cusHandl.service.RoleNewAdmin(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl RoleAdminHandler) RoleAdminUpdateInfoHandler(ctx *fiber.Ctx) error {

	var request admin.AdminRoles

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	cu := admin.AdminRoles{
		AdminRoles: request.AdminRoles,
	}

	service, err := cusHandl.service.RoleAdminUpdateInfoService(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})

}
