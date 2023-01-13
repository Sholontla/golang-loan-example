package link

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/link"
	"github.com/projects/loans/domain/link/service"
	"github.com/projects/loans/middleware"
	utilerrors "github.com/projects/loans/utils/util_errors"
)

type LinkServiceHandler struct {
	service service.ILinkService
}

func (s LinkServiceHandler) GetLinkHandler(ctx *fiber.Ctx) error {
	var request link.Link

	if err := ctx.BodyParser(&request); err != nil {
		jsonErr := utilerrors.NewBadRequestError("Invalid Json Body")
		ctx.JSON(fiber.Map{strconv.Itoa(jsonErr.Status()): jsonErr})
	}

	payLoad, err := middleware.GetUserLogin(ctx)
	if err != nil {
		return err
	}

	l := link.Link{
		LinkCode:   request.LinkCode,
		CustomerId: payLoad,
	}

	s.service.GetLinkService(l)

	return nil
}
