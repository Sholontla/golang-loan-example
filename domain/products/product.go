package products

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/products/dto"
)

type Product struct {
	ProductId           uuid.UUID `db:"product_id"`
	ProductTitle        string    `db:"product_title"`
	ProductCategory     string    `db:"product_category"`
	ProductDescription  string    `db:"product_description"`
	ProductImage        string    `db:"product_image"`
	ProductPrice        float64   `db:"product_price"`
	ProductQuantity     int64     `db:"product_quantity"`
	ProductSerialNumber string    `db:"product_serial_number"`
	ProductCreatedAt    string    `db:"product_created_at"`
	ProductUpdatedAt    string    `db:"product_updated_at"`
	SupplierName        string    `db:"supplier_name"`
}

func (a Product) ToNewProductResponseDto() []dto.ProductResponse {
	response := []dto.ProductResponse{
		{
			ProductId:           a.ProductId,
			ProductTitle:        a.ProductTitle,
			ProductCategory:     a.ProductCategory,
			ProductDescription:  a.ProductDescription,
			ProductImage:        a.ProductImage,
			ProductPrice:        a.ProductPrice,
			ProductQuantity:     a.ProductQuantity,
			ProductSerialNumber: a.ProductSerialNumber,
			ProductCreatedAt:    a.ProductCreatedAt,
			ProductUpdatedAt:    a.ProductUpdatedAt,
			SupplierName:        a.SupplierName,
		},
	}
	return response
}

type ProductFilter struct {
	ProductTitle    string  `db:"product_title"`
	ProductCategory string  `db:"product_category"`
	Min             float64 `db:"min"`
	Max             float64 `db:"max"`
	SupplierName    string  `db:"supplier_name"`
}
