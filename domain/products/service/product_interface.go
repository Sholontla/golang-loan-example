package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/products"
	dto "github.com/projects/loans/domain/products/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type ProductService interface {
	NewProduct(p products.Product) (*products.Product, utilerrors.RestErr)
	FilterProductService(title string, category string, min float64, max float64) ([]dto.ProductResponse, utilerrors.RestErr)
	GetProductService(title string) (*[]dto.ProductResponse, utilerrors.RestErr)
	ProductGetAllService(ctx *fiber.Ctx) ([]dto.ProductResponse, products.PageInfo, utilerrors.RestErr)
	ProductUpdateInfoService(products.Product) (*products.Product, utilerrors.RestErr)
	ProductChacheGetAll(ctx *fiber.Ctx) ([]dto.ProductResponse, utilerrors.RestErr)
	ProductChacheFilterService(title string, category string, min float64, max float64) ([]dto.ProductResponse, utilerrors.RestErr)
}
