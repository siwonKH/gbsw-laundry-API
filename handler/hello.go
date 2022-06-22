package handler

import "github.com/gofiber/fiber/v2"

func Hello(c *fiber.Ctx) error {
	return c.SendString("경북소프트웨어고등학교 세탁실 이용 API")
}
