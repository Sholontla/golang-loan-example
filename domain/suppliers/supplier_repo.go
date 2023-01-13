package suppliers

import utilerrors "github.com/projects/loans/utils/util_errors"

type SupplierRepository interface {
	SuppliersRegistration(Supplier) (*Supplier, *utilerrors.RestErr)
	SuppliersUpdateInfoDao(Supplier) (*Supplier, error)
}
