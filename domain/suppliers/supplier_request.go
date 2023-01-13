package suppliers

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/products"
)

type SupplierRequest struct {
	SupplierId                 uuid.UUID          `db:"supplier_id"`
	SupplierName               string             `db:"supplier_name"`
	SupplierPhoneNumber        string             `db:"supplier_phone_number"`
	SupplierPhoneNumberTwo     string             `db:"supplier_phone_number_two"`
	SupplierPhoneNumberThree   string             `db:"supplier_phone_number_three"`
	SupplierPhoneNumberMessage string             `db:"supplier_phone_number_message"`
	SupplierEmailNumber        string             `db:"supplier_email_number"`
	SupplierEmailNumberTwo     string             `db:"supplier_email_number_two"`
	SupplierEmailNumberThree   string             `db:"supplier_email_number_three"`
	SupplierCreatedAt          string             `db:"supplier_created_at"`
	SupplierUpdatedAt          string             `db:"supplier_updated_at"`
	AdminUser                  string             `db:"admin_user"`
	SupplierProducts           []products.Product `db:"supplier_products"`
}
