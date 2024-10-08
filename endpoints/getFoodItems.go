package endpoints

import (
	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetFoodItems(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	db, _ := DbInterface.GetDbData()

	foodItems := []types.FoodItems{}
	db.Find(&foodItems)

	return c.JSON(fiber.Map{
		"food_items": foodItems,
	})
}
