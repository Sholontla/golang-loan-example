package customers

import (
	"log"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"

	utilerrors "github.com/projects/loans/utils/util_errors"
)

type CustomerRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryInsertCustomer = `INSERT INTO customers(customer_id, customer_user_name, password, first_name, last_name, email, phone_number, customer_created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8);`

	queryFindByEmail = `SELECT customer_id, customer_user_name, password, first_name, last_name, email, phone_number, customer_created_at FROM customers WHERE email=$1;`

	queryUpdateCustomer = `UPDATE customers SET customer_user_name=$1, first_name=$2, last_name=$3, phone_number=$4, customer_updated_at=$5 WHERE email=$6;`

	queryUpdateCustomerPassword = `UPDATE customers SET password=$1 WHERE email=$2;`

	// queryDeleteCustomer   = `DELETE FROM user_name WHERE email=$1;`
)

func (db CustomerRepositoryDb) CustomerRegistration(u Customer) (*Customer, *utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryInsertCustomer)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Customer")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	u.UserValidation()
	stmt.Exec(u.CustomerID, u.CustomerUserName, u.Password, u.FirstName, u.LastName, u.Email, u.PhoneNumber, u.CustomerCreatedt)

	defer stmt.Close()
	return &u, nil
}

func (db CustomerRepositoryDb) CustomerLogin(u Customer) (*Customer, utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryFindByEmail)
	if err != nil {
		utilerrors.NewBadRequestError("Error while getting Customer")
	}
	defer stmt.Close()
	authPass := u.Password

	r := stmt.QueryRow(u.Email)
	r.Scan(&u.CustomerID, &u.CustomerUserName, &u.Password, &u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber, &u.CustomerCreatedt)

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(authPass)); err != nil {
		return nil, utilerrors.NewInvalidPasswordError("Invalid Password")
	}

	return &u, nil
}

func (db CustomerRepositoryDb) CustomerLoggged(payload string) (*Customer, error) {
	stmt, err := db.Client.Prepare(queryFindByEmail)
	if err != nil {
		utilerrors.NewBadRequestError("Error while getting Customer Email")
	}

	var u Customer
	r := stmt.QueryRow(payload)
	r.Scan(&u.CustomerID, &u.CustomerUserName, &u.Password, &u.FirstName, &u.LastName, &u.Email, &u.PhoneNumber, &u.CustomerCreatedt)

	defer stmt.Close()
	return &u, nil
}

func (db CustomerRepositoryDb) CustomerUpdateInfoDao(u Customer) (*Customer, error) {

	stmt, err := db.Client.Prepare(queryUpdateCustomer)
	if err != nil {
		panic(err)
	}

	stmt.Exec(u.CustomerUserName, u.FirstName, u.LastName, u.PhoneNumber, u.CustomerUpdatedAt, u.Email)
	defer stmt.Close()
	return &u, nil

}

func (db CustomerRepositoryDb) CustomerUpdatePasswordDao(u Customer) (*Customer, error) {

	stmt, err := db.Client.Prepare(queryUpdateCustomerPassword)
	if err != nil {
		panic(err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}
	u.Password = string(hash)
	stmt.Exec(u.Password, u.Email)
	defer stmt.Close()
	return &u, nil

}

func NewCustomerRepositoryDb(dBClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dBClient}
}
