package dto

import (
	"github.com/google/uuid"
)

type SupplierResponse struct {
	SupplierId                 uuid.UUID `json:"supplier_id"`
	SupplierName               string    `json:"supplier_name"`
	SupplierPhoneNumber        string    `json:"supplier_phone_number"`
	SupplierPhoneNumberTwo     string    `json:"supplier_phone_number_two"`
	SupplierPhoneNumberThree   string    `json:"supplier_phone_number_three"`
	SupplierPhoneNumberMessage string    `json:"supplier_phone_number_message"`
	SupplieremailNumber        string    `json:"supplier_email_number"`
	SupplieremailNumberTwo     string    `json:"supplier_email_number_two"`
	SupplieremailNumberThree   string    `json:"supplier_email_number_three"`
	SuppliercreatedAt          string    `json:"supplier_created_at"`
	SupplierupdatedAt          string    `json:"supplier_updated_at"`
	AdminUser                  string    `json:"admin_user"`
}
