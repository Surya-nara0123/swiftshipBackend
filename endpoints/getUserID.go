package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetUserID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	// var cookies map[string]string
	// err1 := c.CookieParser(&cookies)
	// if err1 != nil {
	// 	return c.Status(400).JSON(fiber.Map{
	// 		"error": "Cannot parse cookies",
	// 	})
	// }
	// fmt.Println(cookies)
	user := new(types.UserIdReq2)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*user)

	db, _ := DbInterface.GetDbData()

	userDetails := new(types.UserDetails)

	db.First(&userDetails, "username = ?", user.Username)

	id := userDetails.UID

	fmt.Println(id)
	return c.JSON(fiber.Map{
		"uid": id,
	})
}
