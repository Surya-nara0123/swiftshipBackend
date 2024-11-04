package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/surya-nara0123/swiftship/types"
)

// ClearCookies is a middleware that clears all cookies
func SetCookie(c *fiber.Ctx) error {
	user := types.UserDetails{}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	// convert into jwt token
	// set cookie
	t := jwt.New(jwt.SigningMethodHS256)
	key := []byte("secret")
	s, err := t.SignedString(key)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not sign token",
		})
	}

	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: s,
	})

	return c.Next()
}
