package main

import (
	"gbsw-laundry-API/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:                 true,
		CaseSensitive:           true,
		StrictRouting:           true,
		ServerHeader:            "Fiber",
		AppName:                 "gbsw-laundry-API",
		EnableTrustedProxyCheck: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "localhost",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Use(logger.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3507"))
}
