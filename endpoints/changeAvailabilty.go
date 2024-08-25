package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func ChangeAvailability(c *fiber.Ctx, dbInterface database.DatabaseStruct) error {
	foodId := new(types.RestaurantAvailability)
	err := c.BodyParser(foodId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := dbInterface.GetDbData()

	foodItem := types.FoodItems{}
	db.First(&foodItem, "uid = ?", foodId.FoodId)

	foodItem.IsAvailable = !foodItem.IsAvailable

	db.Save(&foodItem)

	return c.JSON(fiber.Map{})

}
