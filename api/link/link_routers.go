package link

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api"
	"github.com/projects/loans/domain/link"
	"github.com/projects/loans/domain/link/service"
	"github.com/projects/loans/middleware"
)

func LinkRouters(r *fiber.App) {

	DbClient := api.GetDbClient()

	accountRepositroyDb := link.NewLinkRepositoryDb(DbClient)

	cHandl := LinkServiceHandler{service.NewLinkService(accountRepositroyDb)}

	api := r.Group("api")
	loan := api.Group("loan")
	admin := loan.Group("link")

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("get", cHandl.GetLinkHandler)

}
