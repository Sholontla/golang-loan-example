package admin

import (
	"github.com/google/uuid"
	"github.com/projects/loans/domain/admin/dto"
)

type Admin struct {
	AdminId          uuid.UUID `db:"admin_id"`
	AdminFirstName   string    `db:"admin_first_name"`
	AdminLastName    string    `db:"admin_last_name"`
	AdminUserName    string    `db:"admin_user_name"`
	AdminEmail       string    `db:"admin_email"`
	AdminPhoneNumber string    `db:"admin_phone_number"`
	AdminAccount     string    `db:"admin_account"`
	AdminPassword    string    `db:"admin_password"`
	AdminCreatedAt   string    `db:"admin_created_at"`
	AdminUpdatedAt   string    `db:"admin_updated_at"`
	IsAdmin          bool      `db:"-"`
}

func (a Admin) ToNewAdminResponseDto() dto.AdminResponse {
	return dto.AdminResponse{
		AdminId:          a.AdminId,
		AdminUserName:    a.AdminUserName,
		AdminPassword:    a.AdminPassword,
		AdminFirstName:   a.AdminFirstName,
		AdminLastName:    a.AdminLastName,
		AdminEmail:       a.AdminEmail,
		AdminPhoneNumber: a.AdminPhoneNumber,
		AdminCreatedAt:   a.AdminCreatedAt,
		AdminUpdatedAt:   a.AdminUpdatedAt,
	}
}
