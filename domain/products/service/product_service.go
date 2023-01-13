package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/projects/loans/domain/products"
	dto "github.com/projects/loans/domain/products/dto"
	"github.com/projects/loans/utils/date_utils"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type ProductRepoService struct {
	repo products.ProductsRepository
}

func (req ProductRepoService) NewProduct(p products.Product) (*products.Product, utilerrors.RestErr) {

	pr := products.Product{
		ProductId:           uuid.New(),
		ProductTitle:        p.ProductTitle,
		ProductCategory:     p.ProductCategory,
		ProductDescription:  p.ProductDescription,
		ProductImage:        p.ProductImage,
		ProductPrice:        p.ProductPrice,
		ProductQuantity:     p.ProductQuantity,
		ProductSerialNumber: "123456789",
		ProductCreatedAt:    date_utils.GetNowString(),
		SupplierName:        p.SupplierName,
	}

	req.repo.ProductRegistration(pr)

	return &pr, nil
}

func (req ProductRepoService) FilterProductService(title string, category string, min float64, max float64) ([]dto.ProductResponse, utilerrors.RestErr) {
	filterProduct, err := req.repo.FilterByProduct(title, category, min, max)
	if err != nil {
		return nil, utilerrors.NewBadRequestError("")
	}
	response := make([]dto.ProductResponse, 0)
	for _, c := range filterProduct {
		response = append(response, c.ToNewProductResponseDto()...)
	}
	return response, nil
}

func (req ProductRepoService) GetProductService(title string) (*[]dto.ProductResponse, utilerrors.RestErr) {

	filterProduct, err := req.repo.GetByProduct(title)
	if err != nil {
		return nil, utilerrors.NewBadRequestError("")
	}
	response := filterProduct.ToNewProductResponseDto()
	return &response, nil
}

func (req ProductRepoService) ProductGetAllService(ctx *fiber.Ctx) ([]dto.ProductResponse, products.PageInfo, utilerrors.RestErr) {
	products, page, err := req.repo.ProductGetAllDao(ctx)
	if err != nil {
		return nil, page, utilerrors.NewBadRequestError("")
	}
	response := make([]dto.ProductResponse, 0)
	for _, c := range products {
		response = append(response, c.ToNewProductResponseDto()...)
	}
	return response, page, nil
}

func (req ProductRepoService) ProductUpdateInfoService(pr products.Product) (*products.Product, utilerrors.RestErr) {

	p := products.Product{
		ProductTitle:       pr.ProductTitle,
		ProductCategory:    pr.ProductCategory,
		ProductDescription: pr.ProductDescription,
		ProductImage:       pr.ProductImage,
		ProductPrice:       pr.ProductPrice,
		ProductUpdatedAt:   date_utils.GetNowString(),
	}

	req.repo.ProductUpdateInfoDao(p)

	return &p, nil
}

func (req ProductRepoService) ProductChacheGetAll(ctx *fiber.Ctx) ([]dto.ProductResponse, utilerrors.RestErr) {

	products, err := req.repo.ProductChacheGetAllDao(ctx)
	if err != nil {
		return nil, utilerrors.NewBadRequestError("")
	}
	response := make([]dto.ProductResponse, 0)
	for _, c := range products {
		response = append(response, c.ToNewProductResponseDto()...)
	}

	return response, nil
}

func (req ProductRepoService) ProductChacheFilterService(title string, category string, min float64, max float64) ([]dto.ProductResponse, utilerrors.RestErr) {
	products, err := req.repo.ProductChacheFilterDao(title, category, min, max)
	if err != nil {
		return nil, utilerrors.NewBadRequestError("")
	}
	response := make([]dto.ProductResponse, 0)
	for _, c := range products {
		response = append(response, c.ToNewProductResponseDto()...)
	}

	return response, nil
}

func NewProductService(repo products.ProductsRepository) ProductRepoService {
	return ProductRepoService{repo: repo}
}
