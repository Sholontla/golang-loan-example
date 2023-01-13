package service

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/products"

	dtoP "github.com/projects/loans/domain/products/dto"
	"github.com/projects/loans/domain/suppliers"
	dto "github.com/projects/loans/domain/suppliers/dto"
	"github.com/projects/loans/utils/date_utils"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type SupplierRepoService struct {
	repo suppliers.SupplierRepository
}

type ProductRepoService struct {
	repo products.ProductsRepository
}

func (req ProductRepoService) NewProduct(p products.Product) (*[]dtoP.ProductResponse, utilerrors.RestErr) {

	newProduct, err := req.repo.ProductRegistration(p)
	if err != nil {
		return nil, utilerrors.NewBadRequestError("")
	}

	response := newProduct.ToNewProductResponseDto()

	return &response, nil
}

func (req SupplierRepoService) NewSupplier(su suppliers.SupplierRequest, payload string) (*dto.SupplierResponse, utilerrors.RestErr) {
	for i := range su.SupplierProducts {

		s := suppliers.Supplier{
			SupplierId:                 uuid.New(),
			SupplierName:               su.SupplierName,
			SupplierPhoneNumber:        su.SupplierPhoneNumber,
			SupplierPhoneNumberTwo:     su.SupplierPhoneNumberTwo,
			SupplierPhoneNumberThree:   su.SupplierPhoneNumberThree,
			SupplierPhoneNumberMessage: su.SupplierPhoneNumberMessage,
			SupplierEmailNumber:        su.SupplierEmailNumber,
			SupplierEmailNumberTwo:     su.SupplierEmailNumberTwo,
			SupplierEmailNumberThree:   su.SupplierEmailNumberThree,
			SupplierCreatedAt:          date_utils.GetNowString(),
			AdminUser:                  payload,
			SupplierProducts: []products.Product{{
				ProductId:           uuid.New(),
				ProductTitle:        su.SupplierProducts[i].ProductTitle,
				ProductCategory:     su.SupplierProducts[i].ProductCategory,
				ProductDescription:  su.SupplierProducts[i].ProductDescription,
				ProductImage:        su.SupplierProducts[i].ProductImage,
				ProductPrice:        su.SupplierProducts[i].ProductPrice,
				ProductQuantity:     su.SupplierProducts[i].ProductQuantity,
				ProductSerialNumber: "123456789",
				ProductCreatedAt:    date_utils.GetNowString(),
				SupplierName:        su.SupplierName},
			},
		}

		newSupplier, err := req.repo.SuppliersRegistration(s)
		if err != nil {
			return nil, utilerrors.NewBadRequestError("")
		}

		response := newSupplier.ToNewAdminResponseDto()

		return &response, nil
	}
	return nil, nil
}

func (req SupplierRepoService) SupplierUpdateInfoService(s suppliers.Supplier) (*suppliers.Supplier, utilerrors.RestErr) {

	su := suppliers.Supplier{
		SupplierPhoneNumber:        s.SupplierPhoneNumber,
		SupplierPhoneNumberTwo:     s.SupplierPhoneNumberTwo,
		SupplierPhoneNumberThree:   s.SupplierPhoneNumberThree,
		SupplierPhoneNumberMessage: s.SupplierPhoneNumberMessage,
		SupplierEmailNumber:        s.SupplierEmailNumber,
		SupplierEmailNumberTwo:     s.SupplierEmailNumberTwo,
		SupplierEmailNumberThree:   s.SupplierEmailNumberThree,
		SupplierUpdatedAt:          date_utils.GetNowString(),
	}

	req.repo.SuppliersUpdateInfoDao(su)

	return &su, nil
}

func NewSupplierService(repo suppliers.SupplierRepository) SupplierRepoService {
	return SupplierRepoService{repo: repo}
}
