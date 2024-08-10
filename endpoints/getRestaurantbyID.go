package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetRestaurantbyID(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	restaurant := new(types.RestuarantIdReq)

	err := c.BodyParser(restaurant)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}
	fmt.Println(*restaurant)

	db, _ := DbInterface.GetDbData()

	restaurantData := types.RestaurantData{}

	db.First(&restaurantData, "uid = ?", restaurant.ID)

	return c.JSON(fiber.Map{
		"result": map[string]interface{}{
			"uid":      restaurantData.UID,
			"name":     restaurantData.Name,
			"location": restaurantData.Location,
			"is_veg":   restaurantData.IsVeg,
		},
		"status": "ok",
	})
}
