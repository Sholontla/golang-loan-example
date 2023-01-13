package dto

import (
	"github.com/google/uuid"
)

type ProductResponse struct {
	ProductId           uuid.UUID `json:"product_id"`
	ProductTitle        string    `json:"product_title"`
	ProductCategory     string    `json:"product_category"`
	ProductDescription  string    `json:"product_description"`
	ProductImage        string    `json:"product_image"`
	ProductPrice        float64   `json:"product_price"`
	ProductQuantity     int64     `json:"product_quantity"`
	ProductSerialNumber string    `json:"product_serial_number"`
	ProductCreatedAt    string    `json:"product_created_at"`
	ProductUpdatedAt    string    `json:"product_updated_at"`
	SupplierName        string    `json:"supplier_name"`
}
