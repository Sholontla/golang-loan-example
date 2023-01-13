package link

import (
	"github.com/jmoiron/sqlx"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type LinkRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryGetLink = `SELECT link_id, link_code, customer_id, link_products FROM link WHERE link_code=$1;`
)

func (db LinkRepositoryDb) GetLinkDao(Link) (*Link, utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryGetLink)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating getting link")
	}

	var l Link

	r := stmt.QueryRow(l.LinkCode)
	r.Scan(&l.LinkId, &l.LinkCode, &l.CustomerId, &l.Products)

	return &l, nil

}

func NewLinkRepositoryDb(dBClient *sqlx.DB) LinkRepositoryDb {
	return LinkRepositoryDb{dBClient}
}
