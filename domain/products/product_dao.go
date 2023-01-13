package products

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	"github.com/projects/loans/datasource/redis"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type ProductRepositoryDb struct {
	Client *sqlx.DB
}

const (
	queryInsertProduct = `INSERT INTO products(product_id, product_title, product_category, product_description, product_image, product_price, product_quantity, product_serial_number, product_created_at, supplier_name) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);`

	queryFindByProduct = `SELECT product_id, product_title, product_category, product_description, product_image, product_price, product_quantity, product_serial_number, product_created_at, supplier_name FROM products WHERE product_title=$1;`

	queryGetAllProduct = `SELECT product_id, product_title, product_category, product_description, product_image, product_price, product_quantity, product_serial_number, product_created_at, supplier_name FROM products;`

	queryFilterProduct = `SELECT product_id, product_title, product_category, product_description, product_image, product_price, product_quantity, product_serial_number, product_created_at, supplier_name FROM products WHERE product_title=$1 AND product_category=$2 AND product_price>=$3 AND product_price<=$4;`

	queryUpdateProduct = `UPDATE products SET product_category=$1, product_description=$2, product_image=$3, product_price=$4, product_updated_at=$5 WHERE product_title=$6;`

	queryDeleteProduct = `DELETE FROM user_name WHERE title=$1;`
)

func (db ProductRepositoryDb) ProductRegistration(p Product) (*Product, *utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryInsertProduct)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Products")
	}
	defer stmt.Close()

	d := Product{
		ProductId:           p.ProductId,
		ProductTitle:        p.ProductTitle,
		ProductCategory:     p.ProductCategory,
		ProductDescription:  p.ProductDescription,
		ProductImage:        p.ProductImage,
		ProductPrice:        p.ProductPrice,
		ProductQuantity:     p.ProductQuantity,
		ProductSerialNumber: p.ProductSerialNumber,
		ProductCreatedAt:    p.ProductCreatedAt,
		SupplierName:        p.SupplierName,
	}

	stmt.Exec(d.ProductId, d.ProductTitle, d.ProductCategory, d.ProductDescription, d.ProductImage, d.ProductPrice, d.ProductQuantity, d.ProductSerialNumber, d.ProductCreatedAt, d.SupplierName)

	return &p, nil
}

func (db ProductRepositoryDb) FilterByProduct(title string, category string, min float64, max float64) ([]Product, utilerrors.RestErr) {
	var products []Product
	stmt, errP := db.Client.Prepare(queryFilterProduct)
	if errP != nil {
		return nil, utilerrors.NewInternalServerError("Error while getting products filter", errP)
	}
	r, err := stmt.Query(title, category, min, max)
	if err != nil {
		return nil, utilerrors.NewInternalServerError("Error while getting products filter", err)
	}

	for r.Next() {

		var p Product
		r.Scan(&p.ProductId, &p.ProductTitle, &p.ProductCategory, &p.ProductDescription, &p.ProductImage, &p.ProductPrice, &p.ProductQuantity, &p.ProductSerialNumber, &p.ProductCreatedAt, &p.SupplierName)

		products = append(products, p)
	}

	return products, nil

}

func (db ProductRepositoryDb) GetByProduct(title string) (*Product, *utilerrors.RestErr) {
	stmt, err := db.Client.Prepare(queryFindByProduct)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Products")
	}

	var p Product
	r := stmt.QueryRow(title)
	r.Scan(&p.ProductId, &p.ProductTitle, &p.ProductCategory, &p.ProductDescription, &p.ProductImage, &p.ProductPrice, &p.ProductQuantity, &p.ProductSerialNumber, &p.ProductCreatedAt, &p.SupplierName)

	defer stmt.Close()
	return &p, nil
}

func (db ProductRepositoryDb) ProductGetAllDao(ctx *fiber.Ctx) ([]Product, PageInfo, utilerrors.RestErr) {

	var products []Product

	db.Client.Select(&products, queryGetAllProduct)

	if sortParam := ctx.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].ProductPrice < products[j].ProductPrice
			})
		} else if sortLower == "desc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].ProductPrice > products[j].ProductPrice
			})
		}
	}
	var totalPage = len(products)

	var data []Product = products

	perPage := len(products) + 1
	page, err := strconv.Atoi(ctx.Query("page", "1"))

	pageInfo := PageInfo{
		TotalData: totalPage,
		Page:      page,
		LastPage:  totalPage/perPage + 1,
	}

	if err != nil {
		return nil, pageInfo, utilerrors.NewBadRequestError("")
	}
	if totalPage <= page*perPage && totalPage >= (page-1)*perPage {
		data = products[(page-1)*perPage : totalPage]
	} else if totalPage >= page*perPage {
		data = products[(page-1)*perPage : page*perPage]
	} else {
		data = []Product{}
	}

	return data, pageInfo, nil
}

func (db ProductRepositoryDb) ProductUpdateInfoDao(p Product) (*Product, error) {

	stmt, err := db.Client.Prepare(queryUpdateProduct)
	if err != nil {
		panic(err)
	}

	stmt.Exec(p.ProductTitle, p.ProductCategory, p.ProductDescription, p.ProductImage, p.ProductPrice, p.ProductUpdatedAt)
	defer stmt.Close()
	return &p, nil

}

func (db ProductRepositoryDb) DeleteProduct(p Product) (*Product, *utilerrors.RestErr) {

	stmt, err := db.Client.Prepare(queryDeleteProduct)
	if err != nil {
		utilerrors.NewBadRequestError("Error while creating new Products")
	}

	stmt.Exec(p.ProductId, p.ProductTitle, p.ProductCategory, p.ProductDescription, p.ProductImage, p.ProductPrice, p.ProductQuantity, p.ProductSerialNumber, p.ProductCreatedAt, p.SupplierName)

	defer stmt.Close()
	return &p, nil
}

func (db ProductRepositoryDb) ProductChacheGetAllDao(ctx *fiber.Ctx) ([]Product, utilerrors.RestErr) {

	var products []Product
	var ctxR = context.Background()

	result, errRedis := redis.Cache.Get(ctxR, "products_forntend").Result()

	if errRedis != nil {

		fmt.Println(errRedis.Error())

		db.Client.Select(&products, queryGetAllProduct)

		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		if err := redis.Cache.Set(ctxR, "products_forntend", bytes, 30*time.Minute).Err(); err != nil {
			panic(err)
		}

	} else {
		json.Unmarshal([]byte(result), &products)
	}

	if sortParam := ctx.Query("sort"); sortParam != "" {
		sortLower := strings.ToLower(sortParam)
		if sortLower == "asc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].ProductPrice < products[j].ProductPrice
			})
		} else if sortLower == "desc" {
			sort.Slice(products, func(i, j int) bool {
				return products[i].ProductPrice > products[j].ProductPrice
			})
		}
	}

	return products, nil
}

func (db ProductRepositoryDb) ProductChacheFilterDao(title string, category string, min float64, max float64) ([]Product, utilerrors.RestErr) {

	var products []Product
	var ctxR = context.Background()

	result, errRedis := redis.Cache.Get(ctxR, "products_filter").Result()

	if errRedis != nil {

		fmt.Println(errRedis.Error())

		db.Client.Select(&products, queryGetAllProduct)

		bytes, err := json.Marshal(products)
		if err != nil {
			panic(err)
		}
		if err := redis.Cache.Set(ctxR, "products_filter", bytes, 30*time.Minute).Err(); err != nil {
			panic(err)
		}

	} else {
		json.Unmarshal([]byte(result), &products)
	}

	stmt, errP := db.Client.Prepare(queryFilterProduct)
	if errP != nil {
		return nil, utilerrors.NewInternalServerError("Error while getting products filter", errP)
	}
	r, err := stmt.Query(title, category, min, max)
	if err != nil {
		return nil, utilerrors.NewInternalServerError("Error while getting products filter", err)
	}

	for r.Next() {

		var p Product
		r.Scan(&p.ProductId, &p.ProductTitle, &p.ProductCategory, &p.ProductDescription, &p.ProductImage, &p.ProductPrice, &p.ProductQuantity, &p.ProductSerialNumber, &p.ProductCreatedAt, &p.SupplierName)

		products = append(products, p)
	}

	return products, nil
}

func NewProductRepositoryDb(dBClient *sqlx.DB) ProductRepositoryDb {
	return ProductRepositoryDb{dBClient}
}
