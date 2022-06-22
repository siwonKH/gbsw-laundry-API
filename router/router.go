package router

import (
	"gbsw-laundry-API/handler"
	"gbsw-laundry-API/handler/status"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	home := app.Group("/")
	home.Get("/", handler.Hello)

	// e.g. /1st floor /female laundry room : /1/f
	room := app.Group("/:floor/:gender")
	room.Get("/", status.RoomStatus)

	// e.g. /1st floor /female laundry room /washer : /1/f/washer
	washers := room.Group("/washer")
	dryers := room.Group("/dryer")

	// e.g. /1st floor /female laundry room /washer /no.2 : 1/f/washer/2
	washer := washers.Group("/:num")
	washer.Get("/", status.WasherStatus)     // 1st floor men washer no.2 status : 1/washer/2
	washer.Post("/", handler.UseWasher)      //    use 1st floor men washer no.2 : (POST) /1/m/washer/2
	washer.Delete("/", handler.DisuseWasher) // disuse 1st floor men washer no.2 : (DELETE) /1/m/washer/2

	dryer := dryers.Group("/:num")
	dryer.Get("/", status.DryerStatus)
	dryer.Post("/", handler.UseDryer)
	dryer.Delete("/", handler.DisuseDryer)
}
