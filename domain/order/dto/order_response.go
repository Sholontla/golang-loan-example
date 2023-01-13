package dto

import (
	"github.com/google/uuid"
)

type OrderResponse struct {
	TransactionId string      `json:"transaction_id"`
	CustomerId    uuid.UUID   `json:"customer_id"`
	Code          string      `json:"code"`
	CustomerEmail string      `json:"customer_email"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	Name          string      `json:"name" `
	Email         string      `json:"email"`
	Address       string      `json:"address"`
	City          string      `json:"city"`
	Country       string      `json:"country"`
	Zip           string      `json:"zip"`
	Complete      bool        `json:"complete"`
	Total         float64     `json:"total"`
	OrderItems    []OrderItem `json:"order_items"`
}

type OrderItem struct {
	OrderId         uuid.UUID `json:"order_item_id"`
	ProductTitle    string    `json:"product_title"`
	Price           float64   `json:"price"`
	Quantity        uint      `json:"quantity"`
	AdminRevenue    float64   `json:"admin_revenue"`
	CustomerRevenue float64   `json:"customer_revenue"`
}
