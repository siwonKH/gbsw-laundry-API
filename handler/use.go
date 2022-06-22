package handler

import (
	"gbsw-laundry-API/config"
	"gbsw-laundry-API/database"
	"gbsw-laundry-API/handler/parser"
	"gbsw-laundry-API/handler/responseModel"
	"gbsw-laundry-API/model"
	"github.com/gofiber/fiber/v2"
	"time"
)

func parseUser(c *fiber.Ctx) (model.User, error) {
	user := model.User{}
	err := c.BodyParser(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func UseWasher(c *fiber.Ctx) error {
	floorNum, machineNum, genderNum, err := parser.Parse3Params(c)
	user, err2 := parseUser(c)
	if err != nil || err2 != nil {
		return c.Status(fiber.StatusNotFound).JSON(responseModel.Status{Success: false, Message: "Not Found"})
	}

	if database.Data[floorNum].Gender[genderNum].Washer[machineNum].IsUsing {
		return c.JSON(responseModel.Status{Success: false, Message: "이미 사용중 입니다"})
	}
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].User = user
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].LastUser = user
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].ExpireTime = time.Now().Add(time.Hour * config.MaxWasherHour)
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].IsUsing = true
	return c.JSON(responseModel.Status{Success: true, Message: "세탁기를 사용합니다"})
}

func UseDryer(c *fiber.Ctx) error {
	floorNum, machineNum, genderNum, err := parser.Parse3Params(c)
	user, err2 := parseUser(c)
	if err != nil || err2 != nil {
		return c.Status(fiber.StatusNotFound).JSON(responseModel.Status{Success: false, Message: "Not Found"})
	}
	if database.Data[floorNum].Gender[genderNum].Dryer[machineNum].IsUsing {
		return c.JSON(responseModel.Status{Success: false, Message: "이 건조기는 이미 사용중 입니다"})
	}

	database.Data[floorNum].Gender[genderNum].Dryer[machineNum].User = user
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].LastUser = user
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].ExpireTime = time.Now().Add(time.Hour * config.MaxDryerHour)
	database.Data[floorNum].Gender[genderNum].Dryer[machineNum].IsUsing = true
	return c.JSON(responseModel.Status{Success: true, Message: "건조기를 사용합니다"})
}
