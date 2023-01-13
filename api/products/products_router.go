package products

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api"
	prodao "github.com/projects/loans/domain/products"

	servicep "github.com/projects/loans/domain/products/service"
	"github.com/projects/loans/middleware"
)

func ProductsRouters(r *fiber.App) {

	DbClient := api.GetDbClient()

	productRepositoryDb := prodao.NewProductRepositoryDb(DbClient)

	cHandlPro := ProductHandler{servicep.NewProductService(productRepositoryDb)}

	api := r.Group("api")
	loan := api.Group("products")
	products := loan.Group("products")
	products.Get("get/all", cHandlPro.GetAllProductHandler)
	products.Get("cache/get/all", cHandlPro.ProductChacheGetAllHandler)
	products.Get("cache/filter", cHandlPro.ProductChacheFilterHandler)
	products.Put("/update", cHandlPro.ProductUpdateInfoHandler)

	productsAuthenticated := products.Use(middleware.IsAuthenticated)
	productsAuthenticated.Post("/register", cHandlPro.ProductRegistrationHandler)
	productsAuthenticated.Get("/filter", cHandlPro.ProductFilterHandler)
	productsAuthenticated.Get("/get", cHandlPro.GetProducHandler)
	productsAuthenticated.Get("/get/all", cHandlPro.GetAllProductHandler)
	productsAuthenticated.Get("/update", cHandlPro.ProductUpdateInfoHandler)

}
