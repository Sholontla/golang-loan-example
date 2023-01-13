package products

import (
	"github.com/gofiber/fiber/v2"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type ProductsRepository interface {
	ProductRegistration(Product) (*Product, *utilerrors.RestErr)
	FilterByProduct(title string, category string, min float64, max float64) ([]Product, utilerrors.RestErr)
	GetByProduct(title string) (*Product, *utilerrors.RestErr)
	ProductGetAllDao(ctx *fiber.Ctx) ([]Product, PageInfo, utilerrors.RestErr)
	ProductUpdateInfoDao(Product) (*Product, error)
	ProductChacheGetAllDao(ctx *fiber.Ctx) ([]Product, utilerrors.RestErr)
	ProductChacheFilterDao(title string, category string, min float64, max float64) ([]Product, utilerrors.RestErr)
}
