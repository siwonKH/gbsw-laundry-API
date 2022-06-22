package handler

import (
	"gbsw-laundry-API/database"
	"gbsw-laundry-API/handler/parser"
	"gbsw-laundry-API/handler/responseModel"
	"gbsw-laundry-API/model"
	"github.com/gofiber/fiber/v2"
)

func DisuseWasher(c *fiber.Ctx) error {
	floorNum, machineNum, genderNum, err := parser.Parse3Params(c)
	user, err2 := parseUser(c)
	if err != nil || err2 != nil {
		return c.Status(fiber.StatusNotFound).JSON(responseModel.Status{Success: false, Message: "Not Found"})
	}
	if !database.Data[floorNum].Gender[genderNum].Washer[machineNum].IsUsing {
		return c.JSON(responseModel.Status{Success: false, Message: "이 세탁기는 사용중이 아닙니다"})
	}
	if database.Data[floorNum].Gender[genderNum].Washer[machineNum].User != user {
		return c.SendString("이 세탁기를 사용중인 사람이 본인이 아닙니다")
	}

	database.Data[floorNum].Gender[genderNum].Washer[machineNum].LastUser = user
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].User = model.User{}
	database.Data[floorNum].Gender[genderNum].Washer[machineNum].IsUsing = false
	return c.JSON(responseModel.Status{Success: true, Message: "세탁기 사용을 해제합니다"})
}

func DisuseDryer(c *fiber.Ctx) error {
	floorNum, machineNum, genderNum, err := parser.Parse3Params(c)
	user, err2 := parseUser(c)
	if err != nil || err2 != nil {
		return c.Status(fiber.StatusNotFound).JSON(responseModel.Status{Success: false, Message: "Not Found"})
	}
	if !database.Data[floorNum].Gender[genderNum].Dryer[machineNum].IsUsing {
		return c.JSON(responseModel.Status{Success: false, Message: "이 건조기는 사용중이 아닙니다"})
	}
	if database.Data[floorNum].Gender[genderNum].Dryer[machineNum].User != user {
		return c.SendString("이 건조기를 사용중인 사람이 본인이 아닙니다")
	}

	database.Data[floorNum].Gender[genderNum].Dryer[machineNum].LastUser = user
	database.Data[floorNum].Gender[genderNum].Dryer[machineNum].User = model.User{}
	database.Data[floorNum].Gender[genderNum].Dryer[machineNum].IsUsing = false
	return c.JSON(responseModel.Status{Success: true, Message: "건조기 사용을 해제합니다"})
}
