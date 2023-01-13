package order

import utilerrors "github.com/projects/loans/utils/util_errors"

type OrderRepository interface {
	GetOrderDao(Order) (*Order, utilerrors.RestErr)
}
