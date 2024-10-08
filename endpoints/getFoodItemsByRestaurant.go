package endpoints

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/surya-nara0123/swiftship/database"
	"github.com/surya-nara0123/swiftship/types"
)

func GetFoodItemsByRestaurant(c *fiber.Ctx, DbInterface database.DatabaseStruct) error {
	foodItem := new(types.FoodItemsRestaurantReq)

	err := c.BodyParser(foodItem)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	db, _ := DbInterface.GetDbData()

	foodItems := []types.FoodItems{}
	db.Find(&foodItems, "restuarant_id = ?", foodItem.RestID)

	fmt.Println(foodItems)

	return c.JSON(fiber.Map{
		"food_items": foodItems,
	})
}
