package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func GetCookies(c *fiber.Ctx) error {
	cookies := c.Cookies("token")
	return c.JSON(fiber.Map{
		"cookies": cookies,
	})
}
