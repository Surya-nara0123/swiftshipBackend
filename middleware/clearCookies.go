package middleware

import (
	"github.com/gofiber/fiber/v2"
)

// ClearCookies is a middleware that clears all cookies
func ClearCookies(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Next()
}
