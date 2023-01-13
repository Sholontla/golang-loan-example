package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api/admin"
	"github.com/projects/loans/api/customers"
	"github.com/projects/loans/api/fanout"
	"github.com/projects/loans/api/link"
	"github.com/projects/loans/api/products"
	"github.com/projects/loans/api/suppliers"
	"github.com/projects/loans/datasource/redis"
)

func ServerStart() {

	app := fiber.New()

	redis.SetUpRedis()

	admin.AdminRouters(app)
	fanout.FanOutRouters(app)
	customers.CustomerRouters(app)
	suppliers.SupplierRouters(app)
	products.ProductsRouters(app)
	link.LinkRouters(app)

	go app.Listen(":3002")

	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Banking/Transaction Server: %v\n", err)
	}
}
