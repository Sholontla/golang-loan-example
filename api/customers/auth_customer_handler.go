package customers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/customers"
	"github.com/projects/loans/domain/customers/service"
	"github.com/projects/loans/middleware"
	utilerrors "github.com/projects/loans/utils/util_errors"
	"github.com/valyala/fasthttp"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (cusHandl CustomerHandler) UserRegistrationHandler(ctx *fiber.Ctx) error {

	var request map[string]string

	//var request customers.CustomerRequest
	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	if request["password"] != request["password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": "password not match ...."})
	}

	cu := customers.Customer{
		CustomerUserName: request["customer_user_name"],
		Password:         request["password"],
		FirstName:        request["first_name"],
		LastName:         request["last_name"],
		Email:            request["customer_email"],
		PhoneNumber:      request["phone_number"],
		IsCustomer:       strings.Contains(ctx.Path(), "/api/loan/customer"),
	}
	service, err := cusHandl.service.NewCustomer(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl CustomerHandler) LoginCustomer(ctx *fiber.Ctx) error {
	var cu customers.CustomerLogin
	if err := ctx.BodyParser(&cu); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	service, err := cusHandl.service.LoginCustomer(ctx, cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}

	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl CustomerHandler) CustomerLoggedHandler(ctx *fiber.Ctx) error {

	payLoad, err := middleware.GetUserLogin(ctx)

	if err != nil {
		return utilerrors.NewMiddleWareError("error in middleWare")
	}

	service, err := cusHandl.service.CustomerLogged(payLoad)
	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(fasthttp.StatusOK).JSON(&fiber.Map{
		"message": service})

}
func (cusHandl CustomerHandler) CustomerLogoutHandler(c *fiber.Ctx) error {

	cusHandl.service.CustomerLogout(c)
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Success Logout"})

}

func (cusHandl CustomerHandler) CustomerUpdateInfoHandler(ctx *fiber.Ctx) error {
	var request customers.CustomerRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}

	cu := customers.Customer{
		CustomerUserName: request.CustomerUserName,
		FirstName:        request.FirstName,
		LastName:         request.LastName,
		Email:            payLoad,
		PhoneNumber:      request.PhoneNumber,
	}

	service, err := cusHandl.service.CustomerUpdateInfoService(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})
}

func (cusHandl CustomerHandler) CustomerUpdatePasswordInfoHandler(ctx *fiber.Ctx) error {
	var request map[string]string

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	if request["password"] != request["password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": "password not match ...."})
	}
	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}
	cu := customers.Customer{
		Password: request["password"],
		Email:    payLoad,
	}

	service, err := cusHandl.service.CustomerUpdatePasswordInfoService(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})
}
