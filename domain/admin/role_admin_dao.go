package admin

import (
	"github.com/jmoiron/sqlx"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type RoleAdminRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryInsertRoleAdmin = `INSERT INTO admin(admin_id, admin_first_name, admin_last_name, admin_user_name, admin_email, admin_phone_number, admin_account, admin_password, admin_created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	// queryFindByRoleAdmin = `SELECT admin_id, admin_first_name, admin_last_name, admin_user_name, admin_email, admin_phone_number, admin_account, admin_password, admin_created_at FROM admin WHERE admin_email=$1;`

	queryUpdateRoleAdmin = `UPDATE admin SET admin_first_name=$1, admin_last_name=$2, admin_user_name=$3, admin_phone_number=$4, admin_account=$5 WHERE admin_email=$6;`

	// queryDeleteCustomer   = `DELETE FROM user_name WHERE email=$1;`

)

func (db RoleAdminRepositoryDb) RoleAdminRegistration(u AdminRoles) (*AdminRoles, *utilerrors.RestErr) {
	stmt, err := db.Client.Prepare(queryInsertRoleAdmin)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Admin")
	}

	stmt.Exec(u.AdminRolesId, u.AdminRoles)

	defer stmt.Close()
	return &u, nil
}

func (db RoleAdminRepositoryDb) RoleAdminUpdateInfoDao(u AdminRoles) (*AdminRoles, error) {

	stmt, err := db.Client.Prepare(queryUpdateRoleAdmin)
	if err != nil {
		panic(err)
	}

	stmt.Exec(u.AdminRolesId, u.AdminRoles)
	defer stmt.Close()
	return &u, nil

}

func RoleNewAdminRepositoryDb(dBClient *sqlx.DB) RoleAdminRepositoryDb {
	return RoleAdminRepositoryDb{dBClient}
}
