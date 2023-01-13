package customers

import (
	"github.com/google/uuid"
)

type CustomerRequest struct {
	CustomerID       uuid.UUID `json:"customer_id"`
	CustomerUserName string    `json:"customer_user_name"`
	Password         string    `json:"password"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Email            string    `json:"email"`
	PhoneNumber      string    `json:"phone_number"`
}
