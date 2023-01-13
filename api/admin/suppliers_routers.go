package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api"
	"github.com/projects/loans/domain/suppliers"
	"github.com/projects/loans/middleware"

	"github.com/projects/loans/domain/suppliers/service"
)

func SupplierRouters(r *fiber.App) {

	DbClient := api.GetDbClient()

	accountRepositroyDb := suppliers.NewSupplierRepositoryDb(DbClient)

	cHandl := SupplierHandler{service.NewSupplierService(accountRepositroyDb)}

	api := r.Group("api")
	loan := api.Group("loan")
	admin := loan.Group("supplier")

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Post("register", cHandl.SupplierRegistrationHandler)

}
