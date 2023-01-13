package suppliers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/products"
	"github.com/projects/loans/domain/suppliers"
	"github.com/projects/loans/domain/suppliers/service"
	"github.com/projects/loans/middleware"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type SupplierHandler struct {
	service service.SupplierService
}

func (cusHandl SupplierHandler) SupplierRegistrationHandler(ctx *fiber.Ctx) error {

	var request suppliers.SupplierRequest

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}
	for i := range request.SupplierProducts {
		s := suppliers.SupplierRequest{
			SupplierName:               request.SupplierName,
			SupplierPhoneNumber:        request.SupplierPhoneNumber,
			SupplierPhoneNumberTwo:     request.SupplierPhoneNumberTwo,
			SupplierPhoneNumberThree:   request.SupplierPhoneNumberThree,
			SupplierPhoneNumberMessage: request.SupplierPhoneNumberMessage,
			SupplierEmailNumber:        request.SupplierEmailNumber,
			SupplierEmailNumberTwo:     request.SupplierEmailNumberTwo,
			SupplierEmailNumberThree:   request.SupplierEmailNumberThree,
			SupplierProducts: []products.Product{

				{
					ProductTitle:       request.SupplierProducts[i].ProductTitle,
					ProductCategory:    request.SupplierProducts[i].ProductCategory,
					ProductDescription: request.SupplierProducts[i].ProductDescription,
					ProductImage:       request.SupplierProducts[i].ProductImage,
					ProductPrice:       request.SupplierProducts[i].ProductPrice,
					ProductQuantity:    request.SupplierProducts[i].ProductQuantity,
					SupplierName:       request.SupplierName,
				},
			},
		}
		payLoad, err := middleware.GetUserLogin(ctx)
		if err != nil {
			return err
		}

		service, err := cusHandl.service.NewSupplier(s, payLoad)

		if err != nil {
			return ctx.JSON(fiber.Map{"Service Error": err})
		}
		return ctx.JSON(fiber.Map{"Supplier": service})
	}
	return nil
}

func (cusHandl SupplierHandler) AdminUpdateInfoHandler(ctx *fiber.Ctx) error {
	var request suppliers.Supplier

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}

	cu := suppliers.Supplier{
		SupplierName:               payLoad,
		SupplierPhoneNumber:        request.SupplierPhoneNumber,
		SupplierPhoneNumberTwo:     request.SupplierPhoneNumberTwo,
		SupplierPhoneNumberThree:   request.SupplierPhoneNumberThree,
		SupplierPhoneNumberMessage: request.SupplierPhoneNumberMessage,
		SupplierEmailNumber:        request.SupplierEmailNumber,
		SupplierEmailNumberTwo:     request.SupplierEmailNumberTwo,
		SupplierEmailNumberThree:   request.SupplierEmailNumberThree,
	}

	service, err := cusHandl.service.SupplierUpdateInfoService(cu)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})

}
