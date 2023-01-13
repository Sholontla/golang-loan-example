package admin

import (
	"github.com/google/uuid"
)

type AdminRequest struct {
	AdminID          uuid.UUID `json:"admin_id"`
	AdminUserName    string    `json:"admin_user_name"`
	AdminPassword    string    `json:"admin_password"`
	AdminFirstName   string    `json:"admin_first_name"`
	AdminLastName    string    `json:"admin_last_name"`
	AdminEmail       string    `json:"admin_email"`
	AdminPhoneNumber string    `json:"admin_phone_number"`
}
