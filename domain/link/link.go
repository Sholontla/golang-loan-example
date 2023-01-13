package link

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/link/dto"
	"github.com/projects/loans/domain/products"
)

type Link struct {
	LinkId     uuid.UUID          `db:"link_id"`
	LinkCode   string             `db:"link_code"`
	CustomerId string             `db:"customer_id"`
	Products   []products.Product `db:"link_products"`
}

func (a Link) ToNewLinkResponseDto() dto.LinkResponse {
	return dto.LinkResponse{
		LinkId:     a.LinkId,
		LinkCode:   a.LinkCode,
		CustomerId: a.CustomerId,
		Products:   a.Products,
	}
}
