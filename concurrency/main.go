package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	Router(app)
	go app.Listen(":1000")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Banking/Transaction Server: %v\n", err)
	}
}
