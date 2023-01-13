package dto

import (
	"github.com/google/uuid"
)

type CustomerResponse struct {
	CustomerID        uuid.UUID `json:"customer_id"`
	CustomerUserName  string    `json:"customer_user_name"`
	Password          string    `json:"password"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"customer_email"`
	PhoneNumber       string    `json:"phone_number"`
	CustomerCreatedt  string    `json:"customer_created_at"`
	CustomerUpdatedAt string    `json:"customer_updated_at"`
}
