package suppliers

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	utilerrors "github.com/projects/loans/utils/util_errors"
)

type SupplierRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryInsertSupplier = `INSERT INTO suppliers(supplier_id, supplier_name, supplier_phone_number, supplier_phone_number_two, supplier_phone_number_three, supplier_phone_number_message, supplier_email_number, supplier_email_number_two, supplier_email_number_three, supplier_created_at, admin_user) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`

	queryInsertProduct = `INSERT INTO products(product_id, product_title, product_category, product_description, product_image, product_price, product_quantity, product_serial_number, product_created_at, supplier_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	// queryFindByEmail = `SELECT admin_id, admin_first_name, admin_last_name, admin_user_name, admin_email, admin_phone_number, admin_account, admin_password, admin_created_at FROM products WHERE title=$1;`

	queryUpdateSupplier = `UPDATE suppliers SET supplier_phone_number=$1, supplier_phone_number_two=$2, supplier_phone_number_three=$3, supplier_phone_number_message=$4, supplier_email_number=$5, supplier_email_number_two=$6, supplier_email_number_three=$7, supplier_updated_at=$8, WHERE supplier_name=$9;`

	// queryDeleteCustomer   = `DELETE FROM user_name WHERE title=$1;`
)

func (db SupplierRepositoryDb) SuppliersRegistration(s Supplier) (*Supplier, *utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryInsertSupplier)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Supplier")
	}
	defer stmt.Close()
	proStmt, err := db.Client.Prepare(queryInsertProduct)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Supplier")
	}
	defer proStmt.Close()
	// proStmt.Exec(s.SupplierProducts.ProductId, s.SupplierProducts.ProductTitle, s.SupplierProducts.ProductCategory, s.SupplierProducts.ProductDescription, s.SupplierProducts.ProductImage, s.SupplierProducts.ProductPrice, s.SupplierProducts.ProductQuantity, s.SupplierProducts.ProductSerialNumber, s.SupplierProducts.ProductCreatedAt, s.SupplierName)
	fmt.Println("DAO: ", s)
	for _, v := range s.SupplierProducts {
		fmt.Println("DAO for loop: ", v)
		proStmt.Exec(v.ProductId, v.ProductTitle, v.ProductCategory, v.ProductDescription, v.ProductImage, v.ProductPrice, v.ProductQuantity, v.ProductSerialNumber, v.ProductCreatedAt, v.SupplierName)
	}

	stmt.Exec(s.SupplierId, s.SupplierName, s.SupplierPhoneNumber, s.SupplierPhoneNumberTwo, s.SupplierPhoneNumberThree, s.SupplierPhoneNumberMessage, s.SupplierEmailNumber, s.SupplierEmailNumberTwo, s.SupplierEmailNumberThree, s.SupplierCreatedAt, s.AdminUser)

	return &s, nil
}

func (db SupplierRepositoryDb) SuppliersUpdateInfoDao(p Supplier) (*Supplier, error) {

	stmt, err := db.Client.Prepare(queryUpdateSupplier)
	if err != nil {
		panic(err)
	}

	stmt.Exec(p.SupplierPhoneNumber, p.SupplierPhoneNumberTwo, p.SupplierPhoneNumberThree, p.SupplierPhoneNumberMessage, p.SupplierEmailNumber, p.SupplierEmailNumberTwo, p.SupplierEmailNumberThree, p.SupplierCreatedAt, p.SupplierUpdatedAt)
	defer stmt.Close()
	return &p, nil

}

func NewSupplierRepositoryDb(dBClient *sqlx.DB) SupplierRepositoryDb {
	return SupplierRepositoryDb{dBClient}
}
