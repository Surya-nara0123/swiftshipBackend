package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetRestaurantID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	restaurant := new(types.RestuarantNameReq)

	err := c.BodyParser(restaurant)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*restaurant)

	db, _ := DbInterface.GetDbData()

	restaurantData := types.RestaurantData{}

	db.First(&restaurantData, "name = ?", restaurant.Name)

	return c.JSON(fiber.Map{
		"result": map[string]interface{}{
			"uid": restaurantData.UID,
		},
		"status": "ok",
	})
}
