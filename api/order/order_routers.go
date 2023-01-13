package link

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api"
	"github.com/projects/loans/domain/order"
	"github.com/projects/loans/domain/order/service"
	"github.com/projects/loans/middleware"
)

func OrderRouters(r *fiber.App) {

	DbClient := api.GetDbClient()

	accountRepositroyDb := order.NewOrderRepositoryDb(DbClient)

	cHandl := OrderServiceHandler{service.NewOrderService(accountRepositroyDb)}

	api := r.Group("api")
	loan := api.Group("loan")
	admin := loan.Group("order")

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("get", cHandl.GetOrderHandler)

}
