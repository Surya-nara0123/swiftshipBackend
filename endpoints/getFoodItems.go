package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetFoodItems(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	db, _ := DbInterface.GetDbData()

	foodItems := []types.FoodItems{}
	db.Find(&foodItems)
	fmt.Println(foodItems[0])

	return c.JSON(fiber.Map{
		"food_items": foodItems,
	})
}
