package fanout

import (
	"github.com/gofiber/fiber/v2"
	"github.com/projects/loans/domain/fanout"
	"github.com/projects/loans/domain/fanout/service"
)

func FanOutHandler(ctx *fiber.Ctx) error {

	var request fanout.Workers
	if err := ctx.BodyParser(&request); err != nil {
		ctx.JSON(fiber.Map{"err": err})
	}

	a := fanout.Worker1{
		Name1:        request.Name1,
		Description1: request.Description1,
		FloatNum1:    request.FloatNum1,
	}

	b := fanout.Worker2{
		Name2:        request.Name2,
		Description2: request.Description2,
		FloatNum2:    request.FloatNum2,
	}

	c := fanout.Worker3{
		Name3:        request.Name3,
		Description3: request.Description3,
		FloatNum3:    request.FloatNum3,
	}

	d := fanout.Worker4{
		Name4:        request.Name4,
		Description4: request.Description4,
		FloatNum4:    request.FloatNum4,
	}

	servicea, serviceb, servicec, serviced, err := service.FanOutService(a, b, c, d)

	if err != nil {
		return ctx.JSON(fiber.Map{"Service Error": err})
	}
	return ctx.JSON(fiber.Map{"WorkerA": servicea, "WorkerB": serviceb, "WorkerC": servicec, "WorkerD": serviced})
}
