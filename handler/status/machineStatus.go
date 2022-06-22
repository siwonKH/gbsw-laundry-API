package status

import (
	"gbsw-laundry-API/config"
	"gbsw-laundry-API/database"
	"gbsw-laundry-API/handler/parser"
	"gbsw-laundry-API/handler/responseModel"
	"github.com/gofiber/fiber/v2"
)

func WasherStatus(c *fiber.Ctx) error {
	floorNum, machineNum, genderNum, err := parser.Parse3Params(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responseModel.Status{Success: false, Message: "400 BadRequest"})
	}
	if 0 > machineNum || machineNum > config.MaxWasherCount {
		return c.Status(fiber.StatusBadRequest).JSON(responseModel.Status{Success: false, Message: "세탁기 번호가 맞지 않습니다"})
	}

	return c.JSON(database.Data[floorNum].Gender[genderNum].Washer[machineNum])
}

func DryerStatus(c *fiber.Ctx) error {
	floorNum, machineNum, genderNum, err := parser.Parse3Params(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responseModel.Status{Success: false, Message: "400 BadRequest"})
	}
	if 0 > machineNum || machineNum > config.MaxDryerCount {
		return c.Status(fiber.StatusBadRequest).JSON(responseModel.Status{Success: false, Message: "건조기 번호가 맞지 않습니다"})
	}

	return c.JSON(database.Data[floorNum].Gender[genderNum].Dryer[machineNum])
}
