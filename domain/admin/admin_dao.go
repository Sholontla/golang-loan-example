package admin

import (
	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	"github.com/projects/loans/domain/security"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type AdminRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryInsertAdmin = `INSERT INTO admin(admin_id, admin_first_name, admin_last_name, admin_user_name, admin_email, admin_phone_number, admin_account, admin_password, admin_created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);`

	queryFindByEmail = `SELECT admin_id, admin_first_name, admin_last_name, admin_user_name, admin_email, admin_phone_number, admin_account, admin_password, admin_created_at FROM admin WHERE admin_email=$1;`

	queryUpdateAdmin = `UPDATE admin SET admin_first_name=$1, admin_last_name=$2, admin_user_name=$3, admin_phone_number=$4, admin_account=$5 WHERE admin_email=$6;`

	queryUpdateAdminPassword = `UPDATE customers SET password=$1 WHERE email=$2;`

	// queryDeleteCustomer   = `DELETE FROM user_name WHERE email=$1;`
)

func (db AdminRepositoryDb) AdminRegistration(u Admin) (*Admin, *utilerrors.RestErr) {
	stmt, err := db.Client.Prepare(queryInsertAdmin)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Admin")
	}
	hash, err := security.SetPassword(u.AdminPassword)
	if err != nil {
		log.Println(err)
	}

	u.AdminPassword = string(hash)
	stmt.Exec(u.AdminId, u.AdminFirstName, u.AdminLastName, u.AdminUserName, u.AdminEmail, u.AdminPhoneNumber, u.AdminAccount, u.AdminPassword, u.AdminCreatedAt)

	defer stmt.Close()
	return &u, nil
}

func (db AdminRepositoryDb) AdminLogin(u Admin) (*Admin, utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryFindByEmail)
	if err != nil {
		utilerrors.NewBadRequestError("Error while getting Admin")
	}
	defer stmt.Close()
	authPass := u.AdminPassword
	r := stmt.QueryRow(u.AdminEmail)
	r.Scan(&u.AdminId, &u.AdminFirstName, &u.AdminLastName, &u.AdminUserName, &u.AdminEmail, &u.AdminPhoneNumber, &u.AdminAccount, &u.AdminPassword, &u.AdminCreatedAt)
	if err := security.ComparePassword(u.AdminPassword, authPass); err != nil {
		return nil, utilerrors.NewInvalidPasswordError("Invalid Password")
	}

	return &u, nil
}

func (db AdminRepositoryDb) AdminLoggged(payload string) (*Admin, error) {
	stmt, err := db.Client.Prepare(queryFindByEmail)
	if err != nil {
		utilerrors.NewBadRequestError("Error while getting Customer Email")
	}

	var u Admin
	r := stmt.QueryRow(payload)
	r.Scan(&u.AdminId, &u.AdminFirstName, &u.AdminLastName, &u.AdminUserName, &u.AdminEmail, &u.AdminPhoneNumber, &u.AdminAccount, &u.AdminPassword, &u.AdminCreatedAt)

	defer stmt.Close()
	return &u, nil
}

func (db AdminRepositoryDb) AdminUpdateInfoDao(u Admin) (*Admin, error) {

	stmt, err := db.Client.Prepare(queryUpdateAdmin)
	if err != nil {
		panic(err)
	}

	stmt.Exec(u.AdminFirstName, u.AdminLastName, u.AdminUserName, u.AdminPhoneNumber, u.AdminUpdatedAt, u.AdminEmail)
	defer stmt.Close()
	return &u, nil

}

func (db AdminRepositoryDb) AdminUpdatePasswordDao(u Admin) (*Admin, error) {

	stmt, err := db.Client.Prepare(queryUpdateAdminPassword)
	if err != nil {
		panic(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.AdminPassword = string(hash)
	stmt.Exec(u.AdminPassword, u.AdminEmail)
	defer stmt.Close()
	return &u, nil

}

func NewAdminRepositoryDb(dBClient *sqlx.DB) AdminRepositoryDb {
	return AdminRepositoryDb{dBClient}
}
