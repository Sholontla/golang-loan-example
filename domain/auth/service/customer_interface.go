package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/customers"
	dto "github.com/projects/loans/domain/customers/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type CustomerService interface {
	NewCustomer(customers.Customer) (*dto.CustomerResponse, *utilerrors.RestErr)
	LoginCustomer(*fiber.Ctx, customers.CustomerLogin) (*dto.CustomerResponse, utilerrors.RestErr)
	CustomerLogged(payload string) (*dto.CustomerResponse, error)
	CustomerLogout(*fiber.Ctx) error
	CustomerUpdateInfoService(customers.Customer) (*customers.Customer, error)
	CustomerUpdatePasswordInfoService(customers.Customer) (*customers.Customer, error)
}
