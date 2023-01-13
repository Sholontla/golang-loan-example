package service

import (
	"github.com/projects/loans/domain/order"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type IOrderService interface {
	GetOrderService(order.Order) (*order.Order, utilerrors.RestErr)
}
