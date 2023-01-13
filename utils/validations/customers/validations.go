package customers

import "github.com/gofiber/fiber/v2"

func PasswordValidation(ctx *fiber.Ctx, request map[string]string) error {
	if request["password"] != request["password_confirm"] {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{"message": "password not patch ...."})
	}
	return nil
}
