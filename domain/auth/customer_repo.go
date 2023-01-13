package customers

import utilerrors "github.com/projects/loans/utils/util_errors"

type CustomerRepository interface {
	CustomerRegistration(Customer) (*Customer, *utilerrors.RestErr)
	CustomerLogin(Customer) (*Customer, utilerrors.RestErr)
	CustomerLoggged(payload string) (*Customer, error)
	CustomerUpdateInfoDao(Customer) (*Customer, error)
	CustomerUpdatePasswordDao(Customer) (*Customer, error)
}
