package customers

import (
	"regexp"
	"strings"

	"github.com/google/uuid"
	"github.com/projects/loans/domain/customers/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type Customer struct {
	CustomerID        uuid.UUID `db:"customer_id"`
	CustomerUserName  string    `db:"customer_user_name"`
	Password          string    `db:"password"`
	FirstName         string    `db:"first_name"`
	LastName          string    `db:"last_name"`
	Email             string    `db:"customer_email"`
	PhoneNumber       string    `db:"phone_number"`
	CustomerCreatedt  string    `db:"customer_created_at"`
	CustomerUpdatedAt string    `db:"customer_updated_at"`
	IsCustomer        bool      `db:"is_customer"`
	Revenue           *float64  `db:"revenue,omitempty"`
}

func (a Customer) ToNewCustomerResponseDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		CustomerID:        a.CustomerID,
		CustomerUserName:  a.CustomerUserName,
		Password:          a.Password,
		FirstName:         a.FirstName,
		LastName:          a.LastName,
		Email:             a.Email,
		PhoneNumber:       a.PhoneNumber,
		CustomerCreatedt:  a.CustomerCreatedt,
		CustomerUpdatedAt: a.CustomerUpdatedAt,
	}
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func (customer *Customer) UserValidation() utilerrors.RestErr {
	customer.Email = strings.TrimSpace(strings.ToLower(customer.Email))

	if customer.Email == "" {
		return utilerrors.NewBadRequestError("Invalid User Email")
	}
	if !isEmailValid(customer.Email) {
		return utilerrors.NewBadRequestError("Invalid User Email")
	}
	if customer.CustomerUserName == "" {
		return utilerrors.NewBadRequestError("Invalid User Name")
	}
	return nil
}
