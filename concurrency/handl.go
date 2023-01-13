package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func ChannelHandl(ctx *fiber.Ctx) error {
	var dataClient Clients
	//Ch := make(chan Clients)

	ctx.BodyParser(&dataClient)

	c, err := ChannelRequestService(dataClient)
	if err != nil {
		log.Println("Error", err.Error())
	}

	return ctx.JSON(c)
}

func NormalHandl(ctx *fiber.Ctx) error {
	var dataClient Clients

	ctx.BodyParser(&dataClient)

	c, err := NormalRequestService(dataClient)
	if err != nil {
		log.Println("Error", err.Error())
	}

	return ctx.JSON(c)
}
