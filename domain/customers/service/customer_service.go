package service

import (
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/projects/loans/domain/customers"
	dto "github.com/projects/loans/domain/customers/dto"
	"github.com/projects/loans/middleware"
	"github.com/projects/loans/utils/date_utils"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type CustomerRepoService struct {
	repo customers.CustomerRepository
}

func (req CustomerRepoService) NewCustomer(u customers.Customer) (*dto.CustomerResponse, *utilerrors.RestErr) {

	c := customers.Customer{
		CustomerID:       uuid.New(),
		CustomerUserName: u.CustomerUserName,
		Password:         u.Password,
		FirstName:        u.FirstName,
		LastName:         u.LastName,
		Email:            u.Email,
		PhoneNumber:      u.PhoneNumber,
		CustomerCreatedt: date_utils.GetNowString(),
		IsCustomer:       u.IsCustomer,
	}

	newCustomer, err := req.repo.CustomerRegistration(c)
	if err != nil {
		return nil, err
	}

	response := newCustomer.ToNewCustomerResponseDto()

	return &response, nil
}

func (req CustomerRepoService) LoginCustomer(ctx *fiber.Ctx, u customers.CustomerLogin) (*dto.CustomerResponse, utilerrors.RestErr) {
	c := customers.Customer{
		Email:    u.Email,
		Password: u.Password,
	}

	newCustomer, err := req.repo.CustomerLogin(c)
	if err != nil {
		return nil, err
	}

	isCustomer := strings.Contains(ctx.Path(), "/api/loan/customer")
	var scope string

	if isCustomer {
		scope = "customer"
	} else {
		scope = "admin"
	}

	nowTime := time.Now()
	expireTime := nowTime.Add(12 * time.Hour)
	token, errL := middleware.GenerateJWT(u.Email, scope)
	if errL != nil {
		ctx.Status(http.StatusOK).JSON(&fiber.Map{
			"message": "Invalid Credentials ..."})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  expireTime,
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	response := newCustomer.ToNewCustomerResponseDto()

	return &response, nil
}

func (req CustomerRepoService) CustomerLogged(payload string) (*dto.CustomerResponse, error) {

	customer, err := req.repo.CustomerLoggged(payload)
	if err != nil {
		return nil, err
	}
	response := customer.ToNewCustomerResponseDto()
	return &response, nil
}

func (cusHandl CustomerRepoService) CustomerLogout(c *fiber.Ctx) error {

	nowTime := time.Now()
	expireTime := nowTime.Add(-time.Hour)
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  expireTime,
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return nil
}

func (req CustomerRepoService) CustomerUpdateInfoService(u customers.Customer) (*customers.Customer, error) {

	c := customers.Customer{
		CustomerUserName:  u.CustomerUserName,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Email:             u.Email,
		PhoneNumber:       u.PhoneNumber,
		CustomerUpdatedAt: date_utils.GetNowString(),
	}

	req.repo.CustomerUpdateInfoDao(c)

	return &c, nil
}

func (req CustomerRepoService) CustomerUpdatePasswordInfoService(u customers.Customer) (*customers.Customer, error) {

	c := customers.Customer{
		Password: u.Password,
		Email:    u.Email,
	}

	req.repo.CustomerUpdatePasswordDao(c)

	return &c, nil
}

func NewCustomerService(repo customers.CustomerRepository) CustomerRepoService {
	return CustomerRepoService{repo: repo}
}
