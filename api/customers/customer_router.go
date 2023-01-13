package customers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api"
	"github.com/projects/loans/domain/customers"
	"github.com/projects/loans/domain/customers/service"
	"github.com/projects/loans/middleware"
)

func CustomerRouters(r *fiber.App) {

	DbClient := api.GetDbClient()

	accountRepositroyDb := customers.NewCustomerRepositoryDb(DbClient)

	cHandl := CustomerHandler{service.NewCustomerService(accountRepositroyDb)}

	api := r.Group("api")
	loan := api.Group("loan")

	customer := loan.Group("customer")
	customer.Post("register", cHandl.UserRegistrationHandler)
	customer.Post("login", cHandl.LoginCustomer)

	customerAuthenticated := customer.Use(middleware.IsAuthenticated)

	customerAuthenticated.Get("/customer", cHandl.CustomerLoggedHandler)
	customerAuthenticated.Put("/update", cHandl.CustomerUpdateInfoHandler)
	customerAuthenticated.Put("/update/password", cHandl.CustomerUpdatePasswordInfoHandler)
	customerAuthenticated.Post("/logout", cHandl.CustomerLogoutHandler)

}
