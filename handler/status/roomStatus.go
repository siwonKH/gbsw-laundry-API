package status

import (
	"gbsw-laundry-API/database"
	"gbsw-laundry-API/handler/parser"
	"gbsw-laundry-API/handler/responseModel"
	"gbsw-laundry-API/model"
	"github.com/gofiber/fiber/v2"
	"time"
)

func RoomStatus(c *fiber.Ctx) error {
	floorNum, genderNum, err := parser.Parse2Params(c)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(responseModel.Status{Success: false, Message: "Not Found"})
	}

	now := time.Now()
	for index, washer := range database.Data[floorNum].Gender[genderNum].Washer {
		if washer.ExpireTime.Sub(now) < 0 {
			washer.LastUser = washer.User
			washer.User = model.User{}
			washer.IsUsing = false
			washer.ExpireTime = now
			database.Data[floorNum].Gender[genderNum].Washer[index] = washer
		}
	}
	for index, dryer := range database.Data[floorNum].Gender[genderNum].Dryer {
		if dryer.ExpireTime.Sub(now) < 0 {
			dryer.LastUser = dryer.User
			dryer.User = model.User{}
			dryer.IsUsing = false
			dryer.ExpireTime = now
			database.Data[floorNum].Gender[genderNum].Dryer[index] = dryer
		}
	}

	return c.JSON(database.Data[floorNum].Gender[genderNum])
}
