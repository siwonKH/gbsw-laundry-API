package parser

import (
	"errors"
	"gbsw-laundry-API/config"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Parse3Params(c *fiber.Ctx) (int, int, int, error) {
	gender := c.Params("gender")
	floorNum, err := strconv.Atoi(c.Params("floor"))
	machineNum, err := strconv.Atoi(c.Params("num"))
	if err != nil {
		return 0, 0, 0, err
	}
	if config.MinFloor > floorNum || floorNum > config.MaxFloor {
		return 0, 0, 0, errors.New("올바른 층이 아닙니다")
	}
	if gender != "m" && gender != "f" {
		return 0, 0, 0, errors.New("올바른 성별이 아닙니다")
	}

	genderNum := 0
	if gender == "f" {
		genderNum = 1
	}
	return floorNum, machineNum, genderNum, nil
}

func Parse2Params(c *fiber.Ctx) (int, int, error) {
	gender := c.Params("gender")
	floorNum, err := strconv.Atoi(c.Params("floor"))
	if err != nil {
		return 0, 0, err
	}
	if config.MinFloor > floorNum || floorNum > config.MaxFloor {
		return 0, 0, errors.New("올바른 층이 아닙니다")
	}
	if gender != "m" && gender != "f" {
		return 0, 0, errors.New("올바른 성별이 아닙니다")
	}

	genderNum := 0
	if gender == "f" {
		genderNum = 1
	}
	return floorNum, genderNum, nil
}
