package link

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/order"
	"github.com/projects/loans/domain/order/service"
	"github.com/projects/loans/middleware"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type OrderServiceHandler struct {
	service service.IOrderService
}

func (s OrderServiceHandler) GetOrderHandler(ctx *fiber.Ctx) error {
	var request order.Order

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}

	o := order.Order{
		Code:          request.Code,
		CustomerEmail: payLoad,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		Name:          request.Name,
		Email:         request.Email,
		Address:       request.Address,
		City:          request.City,
		Country:       request.Country,
		Zip:           request.Zip,
		Complete:      request.Complete,
		Total:         request.Total,
		OrderItems:    request.OrderItems,
	}

	s.service.GetOrderService(o)

	return nil
}
