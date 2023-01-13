package order

import (
	"github.com/jmoiron/sqlx"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type OrderRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryGetorder = `SELECT transaction_id, customer_id, code, customer_email, first_name, last_name, name, email, address, city, country, zip, complete, total, order_items WHERE customer_email=$1;`
)

func (db OrderRepositoryDb) GetOrderDao(Order) (*Order, utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryGetorder)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating getting link")
	}

	var o Order

	r := stmt.QueryRow(o.CustomerEmail)
	r.Scan(&o.TransactionId, &o.CustomerId, &o.Code, &o.CustomerEmail, &o.FirstName, &o.LastName, &o.Name, &o.Email, &o.Address, &o.City, &o.Country, &o.Zip, &o.Complete, &o.Total, &o.OrderItems)

	return &o, nil

}

func NewOrderRepositoryDb(dBClient *sqlx.DB) OrderRepositoryDb {
	return OrderRepositoryDb{dBClient}
}
