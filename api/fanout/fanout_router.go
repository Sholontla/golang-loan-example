package fanout

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/middleware"
)

func FanOutRouters(r *fiber.App) {

	api := r.Group("api")
	loan := api.Group("loan")
	fanout := loan.Group("fanout")

	fanoutAuthenticated := fanout.Use(middleware.IsAuthenticated)

	fanoutAuthenticated.Get("/activate", FanOutHandler)

}
