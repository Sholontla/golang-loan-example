package dto

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/products"
)

type LinkResponse struct {
	LinkId     uuid.UUID          `json:"link_id"`
	LinkCode   string             `json:"link_code"`
	CustomerId string             `json:"customer_id"`
	Products   []products.Product `json:"link_products"`
}
