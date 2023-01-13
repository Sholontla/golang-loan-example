package main

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {

	api := app.Group("api")
	time := api.Group("register")
	time.Post("channel", ChannelHandl)
	time.Post("normal", NormalHandl)

}
