package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/api"
	"github.com/projects/loans/domain/admin"
	supDao "github.com/projects/loans/domain/suppliers"
	supServ "github.com/projects/loans/domain/suppliers/service"

	"github.com/projects/loans/domain/admin/service"
	"github.com/projects/loans/middleware"
)

func AdminRouters(r *fiber.App) {

	DbClient := api.GetDbClient()

	roleAccountRepositroyDb := admin.RoleNewAdminRepositoryDb(DbClient)
	accountRepositroyDb := admin.NewAdminRepositoryDb(DbClient)
	supplierRepositroyDb := supDao.NewSupplierRepositoryDb(DbClient)

	suCHandl := SupplierHandler{supServ.NewSupplierService(supplierRepositroyDb)}

	cHandl := AdminHandler{service.NewAdminService(accountRepositroyDb)}

	roleHandl := RoleAdminHandler{service.NewRoleAdminService(roleAccountRepositroyDb)}

	api := r.Group("api")
	loan := api.Group("loan")
	admin := loan.Group("admin")

	admin.Post("register", cHandl.AdiminRegistrationHandler)
	admin.Post("login", cHandl.LoginAdmin)

	adminAuthenticated := admin.Use(middleware.IsAuthenticated)
	adminAuthenticated.Get("/admin", cHandl.AdminLoggedHandler)
	adminAuthenticated.Put("/update", cHandl.AdminUpdateInfoHandler)
	adminAuthenticated.Put("/update/password", cHandl.AdminUpdatePasswordInfoHandler)
	adminAuthenticated.Post("/logout", cHandl.AdminLogoutHandler)

	adminAuthenticated.Post("role/register", roleHandl.RoleAdiminRegistrationHandler)
	adminAuthenticated.Get("role/get", roleHandl.RoleAdiminRegistrationHandler)
	adminAuthenticated.Put("role/update", roleHandl.RoleAdminUpdateInfoHandler)
	adminAuthenticated.Delete("role/delete", roleHandl.RoleAdiminRegistrationHandler)

	adminAuthenticated.Post("supplier/register", suCHandl.SupplierRegistrationHandler)

}
