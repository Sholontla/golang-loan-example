package service

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/order"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type OrderRepoService struct {
	repo order.OrderRepository
}

func (req OrderRepoService) GetOrderService(or order.Order) (*order.Order, utilerrors.RestErr) {

	o := order.Order{
		TransactionId: uuid.New(),
		CustomerId:    or.CustomerId,
		Code:          or.Code,
		CustomerEmail: or.CustomerEmail,
		FirstName:     or.FirstName,
		LastName:      or.LastName,
		Name:          or.Name,
		Email:         or.Email,
		Address:       or.Address,
		City:          or.City,
		Country:       or.Country,
		Zip:           or.Zip,
		Complete:      or.Complete,
		Total:         or.Total,
		OrderItems:    or.OrderItems,
	}

	req.repo.GetOrderDao(o)

	return &o, nil
}

func NewOrderService(repo order.OrderRepository) OrderRepoService {
	return OrderRepoService{repo: repo}
}
