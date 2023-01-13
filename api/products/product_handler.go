package products

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/products"
	"github.com/projects/loans/domain/products/service"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type ProductHandler struct {
	service service.ProductService
}

func (cusHandl ProductHandler) ProductRegistrationHandler(ctx *fiber.Ctx) error {

	var request products.Product

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}
	p := products.Product{
		ProductTitle:        request.ProductTitle,
		ProductCategory:     request.ProductCategory,
		ProductDescription:  request.ProductDescription,
		ProductImage:        request.ProductImage,
		ProductPrice:        request.ProductPrice,
		ProductQuantity:     request.ProductQuantity,
		ProductSerialNumber: request.ProductSerialNumber,
		ProductCreatedAt:    request.ProductCreatedAt,
		ProductUpdatedAt:    request.ProductUpdatedAt,
		SupplierName:        request.SupplierName,
	}

	service, err := cusHandl.service.NewProduct(p)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl ProductHandler) ProductFilterHandler(ctx *fiber.Ctx) error {

	var request products.ProductFilter

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}
	pF := products.ProductFilter{
		ProductTitle:    request.ProductTitle,
		ProductCategory: request.ProductCategory,
		Min:             request.Min,
		Max:             request.Max,
	}

	service, err := cusHandl.service.FilterProductService(pF.ProductTitle, pF.ProductCategory, pF.Min, pF.Max)
	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl ProductHandler) GetProducHandler(ctx *fiber.Ctx) error {

	var request products.Product

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}
	p := products.Product{
		ProductTitle: request.ProductTitle,
	}

	service, err := cusHandl.service.GetProductService(p.ProductTitle)
	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"message": service})

}

func (cusHandl ProductHandler) GetAllProductHandler(ctx *fiber.Ctx) error {

	service, pageInfo, err := cusHandl.service.ProductGetAllService(ctx)
	if err != nil {
		ctx.JSON(fiber.Map{"Service Error": err})
	}

	return ctx.JSON(fiber.Map{"data": service, "page": pageInfo})
}

func (cusHandl *ProductHandler) ProductUpdateInfoHandler(ctx *fiber.Ctx) error {
	var request products.ProductRequest

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	p := products.Product{
		ProductTitle:       request.ProductTitle,
		ProductCategory:    request.ProductCategory,
		ProductDescription: request.ProductDescription,
		ProductImage:       request.ProductImage,
		ProductPrice:       request.ProductPrice,
		ProductUpdatedAt:   request.ProductUpdatedAt,
	}

	fmt.Println(p)
	service, err := cusHandl.service.ProductUpdateInfoService(p)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": service,
	})

}

func (cusHandl *ProductHandler) ProductChacheGetAllHandler(ctx *fiber.Ctx) error {

	service, err := cusHandl.service.ProductChacheGetAll(ctx)
	if err != nil {
		ctx.JSON(fiber.Map{"Service Error": err})
	}

	return ctx.JSON(service)
}

func (cusHandl *ProductHandler) ProductChacheFilterHandler(ctx *fiber.Ctx) error {
	var request products.ProductFilter

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}
	pF := products.ProductFilter{
		ProductTitle:    request.ProductTitle,
		ProductCategory: request.ProductCategory,
		Min:             request.Min,
		Max:             request.Max,
	}
	service, err := cusHandl.service.ProductChacheFilterService(pF.ProductTitle, pF.ProductCategory, pF.Min, pF.Max)
	if err != nil {
		ctx.JSON(fiber.Map{"Service Error": err})
	}

	return ctx.JSON(service)
}
