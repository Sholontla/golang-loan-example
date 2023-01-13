package service

import (
	"github.com/projects/loans/domain/suppliers"
	dto "github.com/projects/loans/domain/suppliers/dto"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type SupplierService interface {
	NewSupplier(s suppliers.SupplierRequest, payload string) (*dto.SupplierResponse, utilerrors.RestErr)
	SupplierUpdateInfoService(suppliers.Supplier) (*suppliers.Supplier, utilerrors.RestErr)
}
